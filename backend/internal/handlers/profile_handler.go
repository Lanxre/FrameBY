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

type ProfileHandler struct {
	profileService *services.ProfileService
	userService    *services.UserService
}

func NewProfileHandler(profileService *services.ProfileService, userService *services.UserService) *ProfileHandler {
	return &ProfileHandler{
		profileService: profileService,
		userService:    userService,
	}
}

func (h *ProfileHandler) GetBrsmProfile(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	profile, err := h.profileService.GetBrsmProfile(c.Request.Context(), userID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения профиля"})
		return
	}
	if profile == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Профиль не найден"})
		return
	}
	c.JSON(http.StatusOK, profile)
}

func (h *ProfileHandler) CreateBrsmProfile(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	uid := userID.(uuid.UUID)

	existing, _ := h.profileService.GetBrsmProfile(c.Request.Context(), uid)
	if existing != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Вы уже отправили данные представителя БРСМ"})
		return
	}

	hasOther, _ := h.hasOtherProfiles(c, uid)
	if hasOther {
		c.JSON(http.StatusConflict, gin.H{"error": "Вы уже зарегистрированы как представитель другой организации"})
		return
	}

	var req dto.BrsmProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}

	err := h.profileService.CreateBrsmProfile(c.Request.Context(), uid, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания профиля"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Профиль успешно создан"})
}

func (h *ProfileHandler) UpdateBrsmProfile(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	var req dto.BrsmProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}

	err := h.profileService.UpdateBrsmProfile(c.Request.Context(), userID.(uuid.UUID), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления профиля"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Профиль обновлен"})
}

func (h *ProfileHandler) DeleteBrsmProfile(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	err := h.profileService.DeleteBrsmProfile(c.Request.Context(), userID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления профиля"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Профиль удален"})
}

func (h *ProfileHandler) GetStudentProfile(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	profile, err := h.profileService.GetStudentProfile(c.Request.Context(), userID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения профиля"})
		return
	}
	if profile == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Профиль не найден"})
		return
	}
	c.JSON(http.StatusOK, profile)
}

func (h *ProfileHandler) CreateStudentProfile(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	uid := userID.(uuid.UUID)

	existing, _ := h.profileService.GetStudentProfile(c.Request.Context(), uid)
	if existing != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Вы уже отправили данные студента"})
		return
	}

	hasOther, _ := h.hasOtherProfiles(c, uid)
	if hasOther {
		c.JSON(http.StatusConflict, gin.H{"error": "Вы уже зарегистрированы как представитель другой организации"})
		return
	}

	var req dto.StudentProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}

	err := h.profileService.CreateStudentProfile(c.Request.Context(), uid, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания профиля"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Профиль успешно создан"})
}

func (h *ProfileHandler) UpdateStudentProfile(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	var req dto.StudentProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}

	err := h.profileService.UpdateStudentProfile(c.Request.Context(), userID.(uuid.UUID), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления профиля"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Профиль обновлен"})
}

func (h *ProfileHandler) DeleteStudentProfile(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	err := h.profileService.DeleteStudentProfile(c.Request.Context(), userID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления профиля"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Профиль удален"})
}

func (h *ProfileHandler) GetUniversityProfile(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	profile, err := h.profileService.GetUniversityProfile(c.Request.Context(), userID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения профиля"})
		return
	}
	if profile == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Профиль не найден"})
		return
	}
	c.JSON(http.StatusOK, profile)
}

func (h *ProfileHandler) CreateUniversityProfile(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	uid := userID.(uuid.UUID)

	existing, _ := h.profileService.GetUniversityProfile(c.Request.Context(), uid)
	if existing != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Вы уже отправили данные представителя университета"})
		return
	}

	hasOther, _ := h.hasOtherProfiles(c, uid)
	if hasOther {
		c.JSON(http.StatusConflict, gin.H{"error": "Вы уже зарегистрированы как представитель другой организации"})
		return
	}

	var req dto.UniversityProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}

	err := h.profileService.CreateUniversityProfile(c.Request.Context(), uid, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания профиля"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Профиль успешно создан"})
}

func (h *ProfileHandler) UpdateUniversityProfile(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	var req dto.UniversityProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}

	err := h.profileService.UpdateUniversityProfile(c.Request.Context(), userID.(uuid.UUID), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления профиля"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Профиль обновлен"})
}

func (h *ProfileHandler) DeleteUniversityProfile(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	err := h.profileService.DeleteUniversityProfile(c.Request.Context(), userID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления профиля"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Профиль удален"})
}

func (h *ProfileHandler) GetCustomerProfile(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	profile, err := h.profileService.GetCustomerProfile(c.Request.Context(), userID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения профиля"})
		return
	}
	if profile == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Профиль не найден"})
		return
	}
	c.JSON(http.StatusOK, profile)
}

func (h *ProfileHandler) CreateCustomerProfile(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	uid := userID.(uuid.UUID)

	existing, _ := h.profileService.GetCustomerProfile(c.Request.Context(), uid)
	if existing != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Вы уже отправили данные представителя организации"})
		return
	}

	hasOther, _ := h.hasOtherProfiles(c, uid)
	if hasOther {
		c.JSON(http.StatusConflict, gin.H{"error": "Вы уже зарегистрированы как представитель другой организации"})
		return
	}

	var req dto.CustomerProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}

	err := h.profileService.CreateCustomerProfile(c.Request.Context(), uid, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания профиля"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Профиль успешно создан"})
}

func (h *ProfileHandler) UpdateCustomerProfile(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	var req dto.CustomerProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}

	err := h.profileService.UpdateCustomerProfile(c.Request.Context(), userID.(uuid.UUID), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления профиля"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Профиль обновлен"})
}

func (h *ProfileHandler) DeleteCustomerProfile(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	err := h.profileService.DeleteCustomerProfile(c.Request.Context(), userID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления профиля"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Профиль удален"})
}

func (h *ProfileHandler) hasOtherProfiles(c *gin.Context, userID uuid.UUID) (bool, error) {
	brsm, err := h.profileService.GetBrsmProfile(c.Request.Context(), userID)
	if err != nil {
		return false, err
	}
	if brsm != nil {
		return true, nil
	}

	student, err := h.profileService.GetStudentProfile(c.Request.Context(), userID)
	if err != nil {
		return false, err
	}
	if student != nil {
		return true, nil
	}

	university, err := h.profileService.GetUniversityProfile(c.Request.Context(), userID)
	if err != nil {
		return false, err
	}
	if university != nil {
		return true, nil
	}

	customer, err := h.profileService.GetCustomerProfile(c.Request.Context(), userID)
	if err != nil {
		return false, err
	}
	if customer != nil {
		return true, nil
	}

	return false, nil
}

func (h *ProfileHandler) GetStudentProfileByUserID(c *gin.Context) {
	userIDStr := c.Query("user_id")
	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id required"})
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id format"})
		return
	}

	profile, err := h.profileService.GetStudentProfile(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения профиля"})
		return
	}
	if profile == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Профиль студента не найден"})
		return
	}

	response := gin.H{
		"full_name":    profile.FullName,
		"specialty":    profile.Specialty,
		"specialty_id": profile.SpecialtyID,
		"grade":        profile.Grade,
		"position":     profile.Position,
		"phone":        profile.Phone,
	}

	if profile.UniversityDepartmentID != nil {
		dept, err := h.profileService.GetUniversityDepartmentInfo(c.Request.Context(), *profile.UniversityDepartmentID)
		if err == nil && dept != nil {
			response["university_name"] = dept.UniversityName
			response["department_name"] = dept.DepartmentName
		}
	}

	c.JSON(http.StatusOK, response)
}

func (h *ProfileHandler) GetAllProfiles(c *gin.Context) {
	profileType := c.Query("profile_type")
	search := c.Query("search")
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	response, err := h.profileService.GetAllProfiles(c.Request.Context(), profileType, search, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *ProfileHandler) GetUniversityStats(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	uid := userID.(uuid.UUID)

	profile, err := h.profileService.GetUniversityProfile(c.Request.Context(), uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения профиля"})
		return
	}
	if profile == nil || profile.UniversityDepartmentID == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Профиль университета не найден"})
		return
	}

	stats, err := h.profileService.GetUniversityStats(c.Request.Context(), *profile.UniversityDepartmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения статистики"})
		return
	}

	c.JSON(http.StatusOK, stats)
}

func (h *ProfileHandler) UpdateSubrole(c *gin.Context) {
	userIDStr := c.Param(string(middleware.UserIDKey))
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID пользователя"})
		return
	}

	var req dto.UpdateSubroleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}

	err = h.profileService.UpdateProfile(c.Request.Context(), userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления профиля: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Профиль успешно обновлён"})
}

func (h *ProfileHandler) DeleteProfile(c *gin.Context) {
	userIDStr := c.Param(string(middleware.UserIDKey))
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID пользователя"})
		return
	}

	err = h.profileService.DeleteProfile(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления профиля: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Профиль успешно удалён"})
}

func (h *ProfileHandler) SearchUsers(c *gin.Context) {
	search := c.Query("search")
	limitStr := c.DefaultQuery("limit", "10")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	response, err := h.profileService.SearchUsers(c.Request.Context(), search, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *ProfileHandler) GetStudentsWithEmployment(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	uid := userID.(uuid.UUID)

	profile, err := h.profileService.GetUniversityProfile(c.Request.Context(), uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения профиля"})
		return
	}
	if profile == nil || profile.UniversityDepartmentID == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Профиль университета не найден"})
		return
	}

	limitStr := c.DefaultQuery("limit", "20")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 20
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	response, err := h.profileService.GetStudentsWithEmployment(c.Request.Context(), *profile.UniversityDepartmentID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
