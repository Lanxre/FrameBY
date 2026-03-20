package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lanxre/frameby/internal/config"
)

type Route interface {
	Register(cfg *config.Config, r *gin.Engine)
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config, rts ...Route) {
	for _, rt := range rts {
		rt.Register(cfg, r)
	}
}