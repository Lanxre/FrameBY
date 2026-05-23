package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lanxre/frameby/internal/config"
	"github.com/lanxre/frameby/internal/handlers"
	"github.com/lanxre/frameby/internal/middleware"
)

type StudentSquadExportRoutes struct {
	Handler *handlers.StudentSquadExportHandler
	AuthMid *middleware.AuthMiddleware
}

func NewStudentSquadExportRoutes(h *handlers.StudentSquadExportHandler, m *middleware.AuthMiddleware) *StudentSquadExportRoutes {
	return &StudentSquadExportRoutes{
		Handler: h,
		AuthMid: m,
	}
}

func (rts *StudentSquadExportRoutes) Register(cfg *config.Config, r *gin.Engine) {
	apiGroup := r.Group("/api")
	apiGroup.Use(rts.AuthMid.Authenticate())
	{
		apiGroup.GET("/student-squads/:id/export/:format", rts.Handler.ExportByID)
		apiGroup.GET("/student-squads/export/:format", rts.Handler.ExportAll)
	}
}
