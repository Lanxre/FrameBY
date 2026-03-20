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
		// fx.WithLogger(func() fxevent.Logger { return fxevent.NopLogger }),
		fx.WithLogger(func(log *slog.Logger) fxevent.Logger {
					return &fxevent.SlogLogger{Logger: log}
			}),

		fx.Provide(
			config.Load,
			logger.New,			
			storage.NewPostgresDB,
			app.NewValidator,
			
			repositories.NewUserRepository,
			
			services.NewTokenService,
			services.NewAuthService,
			services.NewUserService,
			
			handlers.NewHealthHandler,
			handlers.NewAuthHandler,

			middleware.NewAuthMiddleware,
			
			app.AsRoute(routes.NewHealthRoutes),
			app.AsRoute(routes.NewAuthRoutes),
			
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