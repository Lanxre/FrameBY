package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lanxre/frameby/internal/config"
	"github.com/lanxre/frameby/internal/handlers"
	"github.com/lanxre/frameby/internal/middleware"
)

type ProfileRoutes struct {
	ProfileHandler *handlers.ProfileHandler
	AuthMid        *middleware.AuthMiddleware
}

func NewProfileRoutes(h *handlers.ProfileHandler, m *middleware.AuthMiddleware) *ProfileRoutes {
	return &ProfileRoutes{
		ProfileHandler: h,
		AuthMid:        m,
	}
}

func (rts *ProfileRoutes) Register(cfg *config.Config, r *gin.Engine) {
	authGroup := r.Group("/api")
	authGroup.Use(rts.AuthMid.Authenticate())
	{
		brsm := authGroup.Group("/profile/brsm")
		{
			brsm.GET("", rts.ProfileHandler.GetBrsmProfile)
			brsm.POST("", rts.ProfileHandler.CreateBrsmProfile)
			brsm.PUT("", rts.ProfileHandler.UpdateBrsmProfile)
			brsm.DELETE("", rts.ProfileHandler.DeleteBrsmProfile)
		}

		student := authGroup.Group("/profile/student")
		{
			student.GET("", rts.ProfileHandler.GetStudentProfile)
			student.POST("", rts.ProfileHandler.CreateStudentProfile)
			student.PUT("", rts.ProfileHandler.UpdateStudentProfile)
			student.DELETE("", rts.ProfileHandler.DeleteStudentProfile)
		}

		university := authGroup.Group("/profile/university")
		{
			university.GET("", rts.ProfileHandler.GetUniversityProfile)
			university.POST("", rts.ProfileHandler.CreateUniversityProfile)
			university.PUT("", rts.ProfileHandler.UpdateUniversityProfile)
			university.DELETE("", rts.ProfileHandler.DeleteUniversityProfile)
		}

		customer := authGroup.Group("/profile/customer")
		{
			customer.GET("", rts.ProfileHandler.GetCustomerProfile)
			customer.POST("", rts.ProfileHandler.CreateCustomerProfile)
			customer.PUT("", rts.ProfileHandler.UpdateCustomerProfile)
			customer.DELETE("", rts.ProfileHandler.DeleteCustomerProfile)
		}

		admin := authGroup.Group("/admin/profiles")
		admin.Use(rts.AuthMid.RequireRole("admin"))
		{
			admin.GET("", rts.ProfileHandler.GetAllProfiles)
			admin.PATCH("/:userId/subrole", rts.ProfileHandler.UpdateSubrole)
		}
	}
}
