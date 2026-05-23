package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lanxre/frameby/internal/config"
	"github.com/lanxre/frameby/internal/handlers"
	"github.com/lanxre/frameby/internal/middleware"
)

type SpecialtyRoutes struct {
	Handler *handlers.SpecialtyHandler
	AuthMid *middleware.AuthMiddleware
}

func NewSpecialtyRoutes(h *handlers.SpecialtyHandler, m *middleware.AuthMiddleware) *SpecialtyRoutes {
	return &SpecialtyRoutes{
		Handler: h,
		AuthMid: m,
	}
}

func (rts *SpecialtyRoutes) Register(cfg *config.Config, r *gin.Engine) {
	apiGroup := r.Group("/api")
	apiGroup.Use(rts.AuthMid.Authenticate())
	{
		apiGroup.GET("/specialties", rts.Handler.GetByDepartment)
	}
}
