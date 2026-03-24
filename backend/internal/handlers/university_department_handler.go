package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
