package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lanxre/frameby/internal/config"
	"github.com/lanxre/frameby/internal/handlers"
	"github.com/lanxre/frameby/internal/middleware"
)

type ChatRoutes struct {
	ChatHandler *handlers.ChatHandler
	AuthMid     *middleware.AuthMiddleware
}

func NewChatRoutes(h *handlers.ChatHandler, m *middleware.AuthMiddleware) *ChatRoutes {
	return &ChatRoutes{
		ChatHandler: h,
		AuthMid:     m,
	}
}

func (rts *ChatRoutes) Register(cfg *config.Config, r *gin.Engine) {
	authGroup := r.Group("/api")
	authGroup.Use(rts.AuthMid.Authenticate())
	{
		authGroup.GET("/chats", rts.ChatHandler.GetChats)
		authGroup.POST("/chats", rts.ChatHandler.CreateChat)
		authGroup.GET("/chats/find", rts.ChatHandler.FindDirectChat)
		authGroup.GET("/chats/:id", rts.ChatHandler.GetChat)
		authGroup.GET("/chats/:id/messages", rts.ChatHandler.GetMessages)
		authGroup.POST("/chats/:id/messages", rts.ChatHandler.SendMessage)
		authGroup.POST("/chats/:id/read", rts.ChatHandler.MarkAsRead)
		authGroup.POST("/chats/:id/members/:userId", rts.ChatHandler.AddMember)
		authGroup.DELETE("/chats/:id/members/:userId", rts.ChatHandler.RemoveMember)
	}

	r.GET("/api/ws/chat", rts.ChatHandler.HandleWebSocket)
}
