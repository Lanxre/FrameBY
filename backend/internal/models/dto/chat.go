package dto

type CreateChatRequest struct {
	Type  string   `json:"type" binding:"required"`
	Name  string   `json:"name"`
	Users []string `json:"users"`
}

type ChatResponse struct {
	ID            string  `json:"id"`
	Type          string  `json:"type"`
	Name          *string `json:"name"`
	Avatar        *string `json:"avatar"`
	LastMessage   *string `json:"last_message,omitempty"`
	LastMessageAt *string `json:"last_message_at,omitempty"`
	UnreadCount   int     `json:"unread_count,omitempty"`
	CreatedAt     string  `json:"created_at"`
}

type MessageResponse struct {
	ID         string         `json:"id"`
	ChatID     string         `json:"chat_id"`
	UserID     string         `json:"user_id"`
	UserName   string         `json:"user_name"`
	UserAvatar *string        `json:"user_avatar"`
	Type       string         `json:"type"`
	Content    string         `json:"content"`
	Metadata   map[string]any `json:"metadata,omitempty"`
	CreatedAt  string         `json:"created_at"`
}

type ChatMessagesResponse struct {
	Messages []MessageResponse `json:"messages"`
	Total    int               `json:"total"`
}

type ChatListResponse struct {
	Chats []ChatResponse `json:"chats"`
	Total int            `json:"total"`
}

type AddMemberRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

type SendMessageRequest struct {
	Type     string         `json:"type" binding:"required"`
	Content  string         `json:"content" binding:"required"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

type WebSocketMessage struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type WSChatMessage struct {
	Type    string          `json:"type"`
	ChatID  string          `json:"chat_id"`
	Message MessageResponse `json:"message"`
}

type WSUploadRequest struct {
	Type        string `json:"type"`
	ChatID      string `json:"chat_id"`
	Content     string `json:"content"`
	ContentType string `json:"content_type"`
}
