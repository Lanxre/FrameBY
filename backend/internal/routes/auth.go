package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lanxre/frameby/internal/config"
	"github.com/lanxre/frameby/internal/handlers"
	"github.com/lanxre/frameby/internal/middleware"
)

type AuthRoutes struct {
	AuthHandler *handlers.AuthHandler
	AuthMid     *middleware.AuthMiddleware
}

func NewAuthRoutes(h *handlers.AuthHandler, m *middleware.AuthMiddleware) *AuthRoutes {
	return &AuthRoutes{
		AuthHandler: h,
		AuthMid:     m,
	}
}

func (rts *AuthRoutes) Register(cfg *config.Config, r *gin.Engine) {
	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("/register", rts.AuthHandler.Register)
		authGroup.POST("/login", rts.AuthHandler.Login)
		authGroup.POST("/logout", rts.AuthHandler.Logout)
		
		authGroup.GET("/check", rts.AuthMid.Authenticate(), rts.AuthHandler.CheckAuth)
	}
}