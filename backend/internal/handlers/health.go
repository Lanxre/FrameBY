package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lanxre/frameby/internal/models/dto"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

var startTime = time.Now()

func (h *HealthHandler) HealthCheckHandler(c *gin.Context) {
	response := dto.HealthResponse{
		Status:    "UP",
		Timestamp: time.Now().Format(time.RFC3339),
		Uptime:    time.Since(startTime).String(),
	}

	c.JSON(http.StatusOK, response)
}