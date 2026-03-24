package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lanxre/frameby/internal/config"
	"github.com/lanxre/frameby/internal/handlers"
	"github.com/lanxre/frameby/internal/middleware"
)

type UserRoutes struct {
	UserHandler *handlers.UserHandler
	AuthMid     *middleware.AuthMiddleware
}

func NewUserRoutes(h *handlers.UserHandler, m *middleware.AuthMiddleware) *UserRoutes {
	return &UserRoutes{
		UserHandler: h,
		AuthMid:     m,
	}
}

func (rts *UserRoutes) Register(cfg *config.Config, r *gin.Engine) {
	authGroup := r.Group("/api")
	{
		authGroup.PATCH("/user", rts.AuthMid.Authenticate(), rts.UserHandler.UpdateProfile)
	}
}