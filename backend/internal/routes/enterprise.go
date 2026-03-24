package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lanxre/frameby/internal/config"
	"github.com/lanxre/frameby/internal/handlers"
	"github.com/lanxre/frameby/internal/middleware"
)

type EnterpriseRoutes struct {
	Handler *handlers.EnterpriseHandler
	AuthMid *middleware.AuthMiddleware
}

func NewEnterpriseRoutes(h *handlers.EnterpriseHandler, m *middleware.AuthMiddleware) *EnterpriseRoutes {
	return &EnterpriseRoutes{
		Handler: h,
		AuthMid: m,
	}
}

func (rts *EnterpriseRoutes) Register(cfg *config.Config, r *gin.Engine) {
	apiGroup := r.Group("/api")
	apiGroup.Use(rts.AuthMid.Authenticate())
	{
		apiGroup.GET("/enterprises", rts.Handler.GetAll)
	}
}
