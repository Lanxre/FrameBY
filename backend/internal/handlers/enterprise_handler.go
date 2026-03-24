package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lanxre/frameby/internal/models/dto"
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

func (h *EnterpriseHandler) Create(c *gin.Context) {
	var req dto.CreateEnterpriseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}

	response, err := h.service.Create(c.Request.Context(), req.Name, req.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания организации"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"enterprise": response})
}
