package handlers

import (
	"mime"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lanxre/frameby/internal/middleware"
	"github.com/lanxre/frameby/internal/services"
)

const (
	PDFFormat = "pdf"
	MarkdownFormat = "markdown"
	HTMLFormat = "html"
)

type StudentSquadExportHandler struct {
	service *services.StudentSquadExportService
}

func NewStudentSquadExportHandler(service *services.StudentSquadExportService) *StudentSquadExportHandler {
	return &StudentSquadExportHandler{service: service}
}

func (h *StudentSquadExportHandler) ExportByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	format := c.Param("format")
	if format == "" {
		format = PDFFormat
	}

	data, filename, err := h.service.ExportSquad(c.Request.Context(), id, format)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка экспорта: " + err.Error()})
		return
	}

	contentType := mime.TypeByExtension("." + format)
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	c.Header("Content-Disposition", "attachment; filename=\""+filename+"\"")
	c.Data(http.StatusOK, contentType, data)
}

func (h *StudentSquadExportHandler) ExportAll(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	uid := userID.(uuid.UUID)

	userRole, _ := c.Get(middleware.UserRoleKey)
	role := userRole.(string)

	format := c.Param("format")
	if format == "" {
		format = PDFFormat
	}

	myOnly := c.DefaultQuery("myOnly", "true") == "true"

	data, filename, err := h.service.ExportAllSquads(c.Request.Context(), uid, role, format, myOnly)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка экспорта: " + err.Error()})
		return
	}

	contentType := mime.TypeByExtension("." + format)
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	c.Header("Content-Disposition", "attachment; filename=\""+filename+"\"")
	c.Data(http.StatusOK, contentType, data)
}
