package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lanxre/frameby/internal/config"
	"github.com/lanxre/frameby/internal/handlers"
	"github.com/lanxre/frameby/internal/middleware"
)

type EmploymentRoutes struct {
	Handler *handlers.EmploymentHandler
	AuthMid *middleware.AuthMiddleware
}

func NewEmploymentRoutes(h *handlers.EmploymentHandler, m *middleware.AuthMiddleware) *EmploymentRoutes {
	return &EmploymentRoutes{
		Handler: h,
		AuthMid: m,
	}
}

func (rts *EmploymentRoutes) Register(cfg *config.Config, r *gin.Engine) {
	apiGroup := r.Group("/api")
	apiGroup.Use(rts.AuthMid.Authenticate())
	{
		apiGroup.GET("/employment-requests", rts.Handler.GetAll)
		apiGroup.GET("/employment-requests/statuses", rts.Handler.GetStatuses)
		apiGroup.GET("/employment-requests/participant-statuses", rts.Handler.GetParticipantStatuses)
		apiGroup.GET("/employment-requests/my-applications", rts.Handler.GetMyApplications)
		apiGroup.GET("/employment-requests/:id", rts.Handler.GetByID)
		apiGroup.GET("/employment-requests/:id/participants", rts.Handler.GetParticipants)

		studentGroup := apiGroup.Group("")
		studentGroup.Use(rts.AuthMid.RequireRole(string(middleware.RoleStudent)))
		{
			studentGroup.POST("/employment-requests/apply", rts.Handler.Apply)
			studentGroup.POST("/employment-requests/cancel", rts.Handler.CancelApplication)
		}

		customerGroup := apiGroup.Group("")
		customerGroup.Use(rts.AuthMid.RequireRole(string(middleware.RoleCustomer)))
		{
			customerGroup.POST("/employment-requests", rts.Handler.Create)
			customerGroup.PUT("/employment-requests/:id", rts.Handler.Update)
			customerGroup.DELETE("/employment-requests/:id", rts.Handler.Delete)
			customerGroup.GET("/employment-requests/my-organization", rts.Handler.GetMyOrganizationRequests)
		}

		universityGroup := apiGroup.Group("")
		universityGroup.Use(rts.AuthMid.RequireRole(string(middleware.RoleUniversity)))
		{
			universityGroup.PUT("/employment-requests/:id/participants/:userId/status", rts.Handler.UpdateParticipantStatus)
			universityGroup.POST("/employment-requests/:id/approve", rts.Handler.Approve)
		}
	}
}
