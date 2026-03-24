package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lanxre/frameby/internal/services"
)

type EnterpriseHandler struct {
	service *services.EnterpriseService
}

func NewEnterpriseHandler(service *services.EnterpriseService) *EnterpriseHandler {
	return &EnterpriseHandler{service: service}
}

func (h *EnterpriseHandler) GetAll(c *gin.Context) {
	response, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения организаций"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"enterprises": response})
}
