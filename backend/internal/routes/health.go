package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lanxre/frameby/internal/config"
	"github.com/lanxre/frameby/internal/handlers"
)

type HealthRoutes struct {
	Handler *handlers.HealthHandler
}

func NewHealthRoutes(h *handlers.HealthHandler) *HealthRoutes {
	return &HealthRoutes{
		Handler: h,
	}
}

func (h *HealthRoutes) Register(cfg *config.Config, r *gin.Engine) {
	r.GET("/api/health", h.Handler.HealthCheckHandler)
}