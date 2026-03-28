package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/lanxre/frameby/internal/middleware"
	"github.com/lanxre/frameby/internal/models/dto"
	"github.com/lanxre/frameby/internal/services"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	UserID uuid.UUID
	Conn   *websocket.Conn
	Send   chan []byte
}

type Hub struct {
	clients     map[*Client]bool
	userClients map[uuid.UUID][]*Client
	chats       map[string]map[*Client]bool
	broadcast   chan *BroadcastMessage
	register    chan *Client
	unregister  chan *Client
	mu          sync.RWMutex
}

type BroadcastMessage struct {
	ChatID  string
	Message []byte
	Exclude *Client
}

func NewHub() *Hub {
	return &Hub{
		clients:     make(map[*Client]bool),
		userClients: make(map[uuid.UUID][]*Client),
		chats:       make(map[string]map[*Client]bool),
		broadcast:   make(chan *BroadcastMessage, 256),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.userClients[client.UserID] = append(h.userClients[client.UserID], client)
			h.mu.Unlock()

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)

				if clients, ok := h.userClients[client.UserID]; ok {
					var newClients []*Client
					for _, c := range clients {
						if c != client {
							newClients = append(newClients, c)
						}
					}
					if len(newClients) == 0 {
						delete(h.userClients, client.UserID)
					} else {
						h.userClients[client.UserID] = newClients
					}
				}

				for chatID, chatClients := range h.chats {
					if _, ok := chatClients[client]; ok {
						delete(h.chats[chatID], client)
					}
				}
			}
			h.mu.Unlock()

		case message := <-h.broadcast:
			h.mu.RLock()
			if clients, ok := h.chats[message.ChatID]; ok {
				for client := range clients {
					if message.Exclude != nil && client == message.Exclude {
						continue
					}
					select {
					case client.Send <- message.Message:
					default:
					}
				}
			}
			h.mu.RUnlock()
		}
	}
}

func (h *Hub) JoinChat(client *Client, chatID string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.chats[chatID]; !ok {
		h.chats[chatID] = make(map[*Client]bool)
	}
	h.chats[chatID][client] = true
}

func (h *Hub) LeaveChat(client *Client, chatID string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if clients, ok := h.chats[chatID]; ok {
		delete(clients, client)
		if len(clients) == 0 {
			delete(h.chats, chatID)
		}
	}
}

func (h *Hub) SendToUser(userID uuid.UUID, message []byte) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if clients, ok := h.userClients[userID]; ok {
		for _, client := range clients {
			select {
			case client.Send <- message:
			default:
			}
		}
	}
}

type ChatHandler struct {
	service      *services.ChatService
	tokenService *services.TokenService
	hub          *Hub
}

func NewChatHandler(service *services.ChatService, tokenService *services.TokenService) *ChatHandler {
	hub := NewHub()
	go hub.Run()

	return &ChatHandler{
		service:      service,
		tokenService: tokenService,
		hub:          hub,
	}
}

func (h *ChatHandler) GetHub() *Hub {
	return h.hub
}

func (h *ChatHandler) HandleWebSocket(c *gin.Context) {
	var uid uuid.UUID

	userID, exists := c.Get(middleware.UserIDKey)
	if exists {
		uid = userID.(uuid.UUID)
	} else {
		tokenString := c.Query("token")
		if tokenString == "" {
			cookie, err := c.Cookie("FRAMEBY_ACCESS_TOKEN")
			if err != nil || cookie == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			tokenString = cookie
		}

		claims, err := h.tokenService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		uid = claims.UserID
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}

	client := &Client{
		UserID: uid,
		Conn:   conn,
		Send:   make(chan []byte, 256),
	}

	h.hub.register <- client

	go h.writePump(client)
	go h.readPump(client)
}

func (h *ChatHandler) readPump(client *Client) {
	defer func() {
		h.hub.unregister <- client
		client.Conn.Close()
	}()

	client.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	client.Conn.SetPongHandler(func(string) error {
		client.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			break
		}

		var wsMsg dto.WebSocketMessage
		if err := json.Unmarshal(message, &wsMsg); err != nil {
			continue
		}

		h.handleWebSocketMessage(client, &wsMsg)
	}
}

func (h *ChatHandler) writePump(client *Client) {
	ticker := time.NewTicker(30 * time.Second)
	defer func() {
		ticker.Stop()
		client.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-client.Send:
			client.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := client.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			client.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := client.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (h *ChatHandler) handleWebSocketMessage(client *Client, msg *dto.WebSocketMessage) {
	ctx := context.Background()

	switch msg.Type {
	case "join_chat":
		if chatID, ok := msg.Payload.(string); ok {
			log.Printf("WS: client %s joining chat %s", client.UserID, chatID)
			h.hub.JoinChat(client, chatID)
			h.sendToClient(client, "joined_chat", map[string]string{"chat_id": chatID})
		}

	case "leave_chat":
		if chatID, ok := msg.Payload.(string); ok {
			h.hub.LeaveChat(client, chatID)
			h.sendToClient(client, "left_chat", map[string]string{"chat_id": chatID})
		}

	case "message":
		payload, ok := msg.Payload.(map[string]interface{})
		if !ok {
			return
		}

		chatID, ok := payload["chat_id"].(string)
		if !ok {
			return
		}

		msgType, _ := payload["type"].(string)
		content, _ := payload["content"].(string)

		req := &dto.SendMessageRequest{
			Type:    msgType,
			Content: content,
		}

		chatUUID, err := uuid.Parse(chatID)
		if err != nil {
			return
		}

		fullMsg, err := h.service.SendMessage(ctx, chatUUID, client.UserID, req)
		if err != nil {
			h.sendToClient(client, "error", map[string]string{"error": err.Error()})
			return
		}

		msgResp := dto.MessageResponse{
			ID:         fullMsg.ID,
			ChatID:     fullMsg.ChatID,
			UserID:     fullMsg.UserID,
			UserName:   fullMsg.UserName,
			UserAvatar: fullMsg.UserAvatar,
			Type:       fullMsg.Content.Type,
			Content:    fullMsg.Content.Content,
			Metadata:   fullMsg.Content.Metadata,
			CreatedAt:  fullMsg.CreatedAt.Format("2006-01-02T15:04:05Z"),
		}

		wsMsg := dto.WSChatMessage{
			Type:    "new_message",
			ChatID:  chatID,
			Message: msgResp,
		}

		msgBytes, _ := json.Marshal(wsMsg)
		log.Printf("WS: broadcasting message to chat %s", chatID)
		h.hub.broadcast <- &BroadcastMessage{
			ChatID:  chatID,
			Message: msgBytes,
		}

	case "typing":
		payload, ok := msg.Payload.(map[string]interface{})
		if !ok {
			return
		}

		chatID, ok := payload["chat_id"].(string)
		if !ok {
			return
		}

		wsMsg := dto.WebSocketMessage{
			Type: "user_typing",
			Payload: map[string]interface{}{
				"chat_id": chatID,
				"user_id": client.UserID.String(),
			},
		}

		msgBytes, _ := json.Marshal(wsMsg)
		h.hub.broadcast <- &BroadcastMessage{
			ChatID:  chatID,
			Message: msgBytes,
			Exclude: client,
		}
	}
}

func (h *ChatHandler) sendToClient(client *Client, msgType string, payload interface{}) {
	msg := dto.WebSocketMessage{
		Type:    msgType,
		Payload: payload,
	}
	msgBytes, _ := json.Marshal(msg)
	client.Send <- msgBytes
}

func (h *ChatHandler) GetChats(c *gin.Context) {
	userID := c.MustGet(middleware.UserIDKey).(uuid.UUID)

	limit := 50
	offset := 0

	response, err := h.service.GetUserChats(c.Request.Context(), userID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *ChatHandler) CreateChat(c *gin.Context) {
	userID := c.MustGet(middleware.UserIDKey).(uuid.UUID)

	var req dto.CreateChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chat, err := h.service.CreateChat(c.Request.Context(), userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":   chat.ID,
		"type": chat.Type,
		"name": chat.Name,
	})
}

func (h *ChatHandler) FindDirectChat(c *gin.Context) {
	userID := c.MustGet(middleware.UserIDKey).(uuid.UUID)
	otherUserIDStr := c.Query("user_id")

	otherUserID, err := uuid.Parse(otherUserIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	chat, err := h.service.FindDirectChat(c.Request.Context(), userID, otherUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if chat == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Chat not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   chat.ID,
		"type": chat.Type,
		"name": chat.Name,
	})
}

func (h *ChatHandler) GetChat(c *gin.Context) {
	userID := c.MustGet(middleware.UserIDKey).(uuid.UUID)
	chatID := c.Param("id")

	chatUUID, err := uuid.Parse(chatID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chat ID"})
		return
	}

	isMember, err := h.service.IsMember(c.Request.Context(), chatUUID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !isMember {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not a member of this chat"})
		return
	}

	chat, err := h.service.GetChatByID(c.Request.Context(), chatUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, chat)
}

func (h *ChatHandler) GetMessages(c *gin.Context) {
	userID := c.MustGet(middleware.UserIDKey).(uuid.UUID)
	chatID := c.Param("id")

	chatUUID, err := uuid.Parse(chatID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chat ID"})
		return
	}

	limit := 50
	offset := 0

	response, err := h.service.GetMessages(c.Request.Context(), chatUUID, userID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *ChatHandler) SendMessage(c *gin.Context) {
	userID := c.MustGet(middleware.UserIDKey).(uuid.UUID)
	chatID := c.Param("id")

	chatUUID, err := uuid.Parse(chatID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chat ID"})
		return
	}

	var req dto.SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg, err := h.service.SendMessage(c.Request.Context(), chatUUID, userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	msgResp := dto.MessageResponse{
		ID:         msg.ID,
		ChatID:     msg.ChatID,
		UserID:     msg.UserID,
		UserName:   msg.UserName,
		UserAvatar: msg.UserAvatar,
		Type:       msg.Content.Type,
		Content:    msg.Content.Content,
		Metadata:   msg.Content.Metadata,
		CreatedAt:  msg.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}

	msgBytes, _ := json.Marshal(dto.WSChatMessage{
		Type:    "new_message",
		ChatID:  chatID,
		Message: msgResp,
	})
	h.hub.broadcast <- &BroadcastMessage{
		ChatID:  chatID,
		Message: msgBytes,
	}

	c.JSON(http.StatusCreated, msgResp)
}

func (h *ChatHandler) MarkAsRead(c *gin.Context) {
	userID := c.MustGet(middleware.UserIDKey).(uuid.UUID)
	chatID := c.Param("id")

	chatUUID, err := uuid.Parse(chatID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chat ID"})
		return
	}

	if err := h.service.MarkAsRead(c.Request.Context(), chatUUID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Chat marked as read"})
}

func (h *ChatHandler) AddMember(c *gin.Context) {
	chatID := c.Param("id")
	userIDStr := c.Param("userId")

	chatUUID, err := uuid.Parse(chatID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chat ID"})
		return
	}

	userUUID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.service.AddMember(c.Request.Context(), chatUUID, userUUID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Member added"})
}

func (h *ChatHandler) RemoveMember(c *gin.Context) {
	chatID := c.Param("id")
	userIDStr := c.Param("userId")

	chatUUID, err := uuid.Parse(chatID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chat ID"})
		return
	}

	userUUID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.service.RemoveMember(c.Request.Context(), chatUUID, userUUID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Member removed"})
}
