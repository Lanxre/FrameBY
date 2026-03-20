package app

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lanxre/frameby/internal/config"
	"github.com/lanxre/frameby/internal/utils/static"
	"go.uber.org/fx"
)

func StartHTTPServer(lc fx.Lifecycle, r *gin.Engine, cfg *config.Config) {	
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Requested-With", "Cookie"},
		AllowCredentials: true,
	}))

	srv := &http.Server{
		Handler:      r,
		Addr:         cfg.Address,
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Printf("Server starting on %s ...", cfg.Address)
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("listen: %s\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Shutting down server...")
			return srv.Shutdown(ctx)
		},
	})
}

func RegisterStaticFiles(r *gin.Engine) {
	static.CreateStaticDirs(r)
}