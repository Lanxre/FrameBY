package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lanxre/frameby/internal/middleware"
	"github.com/lanxre/frameby/internal/models/dto"
	"github.com/lanxre/frameby/internal/services"
)

type StudentSquadHandler struct {
	service *services.StudentSquadService
}

func NewStudentSquadHandler(service *services.StudentSquadService) *StudentSquadHandler {
	return &StudentSquadHandler{
		service: service,
	}
}

func (h *StudentSquadHandler) GetAll(c *gin.Context) {
	status := c.Query("status")
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, _ := strconv.Atoi(limitStr)
	if limit <= 0 {
		limit = 10
	}

	offset, _ := strconv.Atoi(offsetStr)
	if offset < 0 {
		offset = 0
	}

	response, err := h.service.GetAll(c.Request.Context(), status, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения отрядов"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *StudentSquadHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	squad, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения отряда"})
		return
	}
	if squad == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Отряд не найден"})
		return
	}

	c.JSON(http.StatusOK, squad)
}

func (h *StudentSquadHandler) Create(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	uid := userID.(uuid.UUID)

	userRole, _ := c.Get(middleware.UserRoleKey)

	var req dto.CreateStudentSquadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}

	if role, ok := userRole.(string); ok {
		req.Role = role
	}

	squad, err := h.service.Create(c.Request.Context(), uid, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"squad": squad})
}

func (h *StudentSquadHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	var req dto.UpdateStudentSquadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}

	if err := h.service.Update(c.Request.Context(), id, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления отряда"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Отряд обновлён"})
}

func (h *StudentSquadHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления отряда"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Отряд удалён"})
}

func (h *StudentSquadHandler) Join(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	uid := userID.(uuid.UUID)

	var req dto.JoinSquadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}

	squadID, err := uuid.Parse(req.SquadID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID отряда"})
		return
	}

	if err := h.service.Join(c.Request.Context(), squadID, uid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Вы вступили в отряд"})
}

func (h *StudentSquadHandler) Leave(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	uid := userID.(uuid.UUID)

	var req dto.JoinSquadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}

	squadID, err := uuid.Parse(req.SquadID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID отряда"})
		return
	}

	if err := h.service.Leave(c.Request.Context(), squadID, uid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка выхода из отряда"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Вы покинули отряд"})
}

func (h *StudentSquadHandler) GetStatuses(c *gin.Context) {
	statuses, err := h.service.GetStatuses(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения статусов"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"statuses": statuses})
}

func (h *StudentSquadHandler) GetMySquads(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	uid := userID.(uuid.UUID)

	userRole, _ := c.Get(middleware.UserRoleKey)
	role := userRole.(string)

	var squads []dto.StudentSquadResponse
	var err error

	switch role {
	case "student":
		squads, err = h.service.GetByParticipant(c.Request.Context(), uid)
	case "university":
		squads, err = h.service.GetByApprover(c.Request.Context(), uid)
	default:
		squads, err = h.service.GetByOrganizer(c.Request.Context(), uid)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения отрядов: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"squads": squads, "role": role, "userId": uid.String()})
}

func (h *StudentSquadHandler) Approve(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	uid := userID.(uuid.UUID)

	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	var req dto.ApproveSquadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	if err := h.service.Approve(c.Request.Context(), id, req.Approved, uid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления статуса"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Статус обновлён"})
}
