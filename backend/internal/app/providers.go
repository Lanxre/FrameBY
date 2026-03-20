package app

import (
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lanxre/frameby/internal/config"
	"github.com/lanxre/frameby/internal/routes"
	"go.uber.org/fx"
)

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(routes.Route)),
		fx.ResultTags(`group:"routes"`),
	)
}

func NewValidator() *validator.Validate {
	return validator.New()
}

func NewGinRouter(cfg *config.Config, rts []routes.Route, log *slog.Logger) *gin.Engine {
	r := gin.New()

	r.Use(func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		log.Info("request completed",
			slog.Int("status", c.Writer.Status()),
			slog.String("method", c.Request.Method),
			slog.String("path", path),
			slog.String("query", query),
			slog.String("ip", c.ClientIP()),
			slog.Duration("duration", time.Since(start)),
		)
	})

	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered any) {
		log.Error("panic recovered",
			slog.Any("error", recovered),
			slog.String("stack", string(debug.Stack())),
		)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
	}))

	routes.RegisterRoutes(r, cfg, rts...)

	return r
}