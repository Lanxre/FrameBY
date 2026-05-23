package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lanxre/frameby/internal/services"
)

type SpecialtyHandler struct {
	service *services.SpecialtyService
}

func NewSpecialtyHandler(service *services.SpecialtyService) *SpecialtyHandler {
	return &SpecialtyHandler{service: service}
}

func (h *SpecialtyHandler) GetByDepartment(c *gin.Context) {
	departmentIDStr := c.Query("department_id")
	universityDeptIDStr := c.Query("university_department_id")

	switch {
	case universityDeptIDStr != "":
		universityDeptID, err := uuid.Parse(universityDeptIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid university_department_id"})
			return
		}
		specialties, err := h.service.GetByUniversityDepartmentID(c.Request.Context(), universityDeptID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения специальностей"})
			return
		}
		c.JSON(http.StatusOK, specialties)

	case departmentIDStr != "":
		departmentID, err := uuid.Parse(departmentIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department_id"})
			return
		}
		specialties, err := h.service.GetByDepartmentID(c.Request.Context(), departmentID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения специальностей"})
			return
		}
		c.JSON(http.StatusOK, specialties)

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "department_id or university_department_id required"})
	}
}
