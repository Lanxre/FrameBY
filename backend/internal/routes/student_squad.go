package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lanxre/frameby/internal/config"
	"github.com/lanxre/frameby/internal/handlers"
	"github.com/lanxre/frameby/internal/middleware"
)

type StudentSquadRoutes struct {
	Handler *handlers.StudentSquadHandler
	AuthMid *middleware.AuthMiddleware
}

func NewStudentSquadRoutes(h *handlers.StudentSquadHandler, m *middleware.AuthMiddleware) *StudentSquadRoutes {
	return &StudentSquadRoutes{
		Handler: h,
		AuthMid: m,
	}
}

func (rts *StudentSquadRoutes) Register(cfg *config.Config, r *gin.Engine) {
	apiGroup := r.Group("/api")
	apiGroup.Use(rts.AuthMid.Authenticate())
	{
		apiGroup.GET("/student-squads", rts.Handler.GetAll)
		apiGroup.GET("/student-squads/statuses", rts.Handler.GetStatuses)
		apiGroup.GET("/student-squads/my", rts.Handler.GetMySquads)
		apiGroup.GET("/student-squads/:id", rts.Handler.GetByID)

		apiGroup.POST("/student-squads/join", rts.Handler.Join)
		apiGroup.POST("/student-squads/leave", rts.Handler.Leave)

		creatorGroup := apiGroup.Group("")
		creatorGroup.Use(rts.AuthMid.RequireAnyRole(string(middleware.RoleBRSM), string(middleware.RoleUniversity), string(middleware.RoleCustomer)))
		{
			creatorGroup.PATCH("/student-squads/:id", rts.Handler.Update)
			creatorGroup.POST("/student-squads", rts.Handler.Create)
		}

		universityGroup := apiGroup.Group("")
		universityGroup.Use(rts.AuthMid.RequireRole(string(middleware.RoleUniversity)))
		{
			universityGroup.POST("/student-squads/:id/approve", rts.Handler.Approve)
		}
	}
}
