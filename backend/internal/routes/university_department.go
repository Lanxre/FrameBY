package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lanxre/frameby/internal/config"
	"github.com/lanxre/frameby/internal/handlers"
	"github.com/lanxre/frameby/internal/middleware"
)

type UniversityDepartmentRoutes struct {
	Handler *handlers.UniversityDepartmentHandler
	AuthMid *middleware.AuthMiddleware
}

func NewUniversityDepartmentRoutes(h *handlers.UniversityDepartmentHandler, m *middleware.AuthMiddleware) *UniversityDepartmentRoutes {
	return &UniversityDepartmentRoutes{
		Handler: h,
		AuthMid: m,
	}
}

func (rts *UniversityDepartmentRoutes) Register(cfg *config.Config, r *gin.Engine) {
	apiGroup := r.Group("/api")
	apiGroup.Use(rts.AuthMid.Authenticate())
	{
		apiGroup.GET("/university-departments", rts.Handler.GetAll)
	}

	adminGroup := r.Group("/api")
	adminGroup.Use(rts.AuthMid.Authenticate())
	adminGroup.Use(rts.AuthMid.RequireRole("admin"))
	{
		adminGroup.POST("/university-departments", rts.Handler.Create)
	}
}
