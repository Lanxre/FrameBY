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

type EmploymentHandler struct {
	service *services.EmploymentService
}

func NewEmploymentHandler(service *services.EmploymentService) *EmploymentHandler {
	return &EmploymentHandler{service: service}
}

func (h *EmploymentHandler) GetAll(c *gin.Context) {
	status := c.Query("status")
	universityDeptID := c.Query("university_department_id")
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

	response, err := h.service.GetAll(c.Request.Context(), status, universityDeptID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения заявок"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *EmploymentHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	request, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения заявки"})
		return
	}
	if request == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заявка не найдена"})
		return
	}

	c.JSON(http.StatusOK, request)
}

func (h *EmploymentHandler) Create(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	uid := userID.(uuid.UUID)

	var req dto.CreateEmploymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}

	universityDeptID, err := uuid.Parse(req.UniversityDepartmentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID университета"})
		return
	}

	request, err := h.service.Create(c.Request.Context(), uid, universityDeptID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания заявки"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"request": request})
}

func (h *EmploymentHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	var req dto.UpdateEmploymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}

	if err := h.service.Update(c.Request.Context(), id, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления заявки"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Заявка обновлена"})
}

func (h *EmploymentHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления заявки"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Заявка удалена"})
}

func (h *EmploymentHandler) Apply(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	uid := userID.(uuid.UUID)

	var req dto.ApplyToEmploymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}

	requestID, err := uuid.Parse(req.RequestID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID заявки"})
		return
	}

	if err := h.service.Apply(c.Request.Context(), requestID, uid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Заявка подана"})
}

func (h *EmploymentHandler) CancelApplication(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	uid := userID.(uuid.UUID)

	var req dto.ApplyToEmploymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}

	requestID, err := uuid.Parse(req.RequestID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID заявки"})
		return
	}

	if err := h.service.CancelApplication(c.Request.Context(), requestID, uid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка отмены заявки"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Заявка отменена"})
}

func (h *EmploymentHandler) GetParticipants(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	participants, err := h.service.GetParticipants(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения участников"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"participants": participants, "total": len(participants)})
}

func (h *EmploymentHandler) UpdateParticipantStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	var req dto.UpdateParticipantStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID пользователя"})
		return
	}

	if err := h.service.UpdateParticipantStatus(c.Request.Context(), id, userID, req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления статуса"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Статус обновлён"})
}

func (h *EmploymentHandler) GetStatuses(c *gin.Context) {
	statuses, err := h.service.GetStatuses(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения статусов"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"statuses": statuses})
}

func (h *EmploymentHandler) GetParticipantStatuses(c *gin.Context) {
	statuses, err := h.service.GetParticipantStatuses(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения статусов"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"statuses": statuses})
}

func (h *EmploymentHandler) GetMyApplications(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	uid := userID.(uuid.UUID)

	requests, err := h.service.GetUserApplications(c.Request.Context(), uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения заявок"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"requests": requests})
}

func (h *EmploymentHandler) GetMyOrganizationRequests(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	uid := userID.(uuid.UUID)

	requests, err := h.service.GetByEnterprise(c.Request.Context(), uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения заявок"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"requests": requests})
}
