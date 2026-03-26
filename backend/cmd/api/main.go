package main

import (
	"log/slog"

	"github.com/lanxre/frameby/internal/app"
	"github.com/lanxre/frameby/internal/config"
	"github.com/lanxre/frameby/internal/handlers"
	"github.com/lanxre/frameby/internal/lib/logger"
	"github.com/lanxre/frameby/internal/middleware"
	"github.com/lanxre/frameby/internal/repositories"
	"github.com/lanxre/frameby/internal/routes"
	"github.com/lanxre/frameby/internal/services"
	"github.com/lanxre/frameby/internal/storage"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	fx.New(
		fx.WithLogger(func(log *slog.Logger) fxevent.Logger {
			return &fxevent.SlogLogger{Logger: log}
		}),

		fx.Provide(
			config.Load,
			logger.New,
			storage.NewPostgresDB,
			app.NewValidator,

			repositories.NewUserRepository,
			repositories.NewProfileRepository,
			repositories.NewUniversityDepartmentRepository,
			repositories.NewEnterpriseRepository,
			repositories.NewStudentSquadRepository,
			repositories.NewEmploymentRepository,

			services.NewTokenService,
			services.NewAuthService,
			services.NewUserService,
			services.NewProfileService,
			services.NewUniversityDepartmentService,
			services.NewEnterpriseService,
			services.NewStudentSquadService,
			services.NewEmploymentService,

			handlers.NewHealthHandler,
			handlers.NewAuthHandler,
			handlers.NewUserHandler,
			handlers.NewProfileHandler,
			handlers.NewUniversityDepartmentHandler,
			handlers.NewEnterpriseHandler,
			handlers.NewStudentSquadHandler,
			handlers.NewEmploymentHandler,

			middleware.NewAuthMiddleware,

			app.AsRoute(routes.NewHealthRoutes),
			app.AsRoute(routes.NewAuthRoutes),
			app.AsRoute(routes.NewUserRoutes),
			app.AsRoute(routes.NewProfileRoutes),
			app.AsRoute(routes.NewUniversityDepartmentRoutes),
			app.AsRoute(routes.NewEnterpriseRoutes),
			app.AsRoute(routes.NewStudentSquadRoutes),
			app.AsRoute(routes.NewEmploymentRoutes),

			fx.Annotate(
				app.NewGinRouter,
				fx.ParamTags(``, `group:"routes"`, ``),
			),
		),

		fx.Invoke(
			app.RegisterStaticFiles,
			app.StartHTTPServer,
		),
	).Run()
}
