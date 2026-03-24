package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lanxre/frameby/internal/middleware"
	"github.com/lanxre/frameby/internal/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID, exists := c.Get(middleware.UserIDKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не авторизован"})
		return
	}

	file, err := c.FormFile("avatar")
	if err == nil && file != nil {
		_, err := h.userService.SaveAvatar(c.Request.Context(), file, userID.(uuid.UUID))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if err := h.userService.UpdateProfile(c.Request.Context(), userID.(uuid.UUID), c.PostForm("login"), c.PostForm("password")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления профиля"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Профиль обновлен",
	})
}
