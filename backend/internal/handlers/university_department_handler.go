package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lanxre/frameby/internal/models/dto"
	"github.com/lanxre/frameby/internal/services"
)

type UniversityDepartmentHandler struct {
	service *services.UniversityDepartmentService
}

func NewUniversityDepartmentHandler(service *services.UniversityDepartmentService) *UniversityDepartmentHandler {
	return &UniversityDepartmentHandler{service: service}
}

func (h *UniversityDepartmentHandler) GetAll(c *gin.Context) {
	response, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения данных"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"departments": response})
}

func (h *UniversityDepartmentHandler) Create(c *gin.Context) {
	var req dto.CreateUniversityDepartmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}

	response, err := h.service.Create(c.Request.Context(), req.UniversityName, req.DepartmentName, req.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания записи"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"department": response})
}
