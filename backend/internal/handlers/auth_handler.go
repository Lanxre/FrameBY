package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lanxre/frameby/internal/config"
	"github.com/lanxre/frameby/internal/models/dto"
	"github.com/lanxre/frameby/internal/services"
)

type AuthHandler struct {
	authSvc  *services.AuthService
	tokenSvc *services.TokenService
	cfg      *config.Config
}

func NewAuthHandler(authSvc *services.AuthService, tokenSvc *services.TokenService, cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		authSvc:  authSvc,
		tokenSvc: tokenSvc,
		cfg:      cfg,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные запроса"})
		return
	}

	err := h.authSvc.Register(c.Request.Context(), req.Email, req.Login, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при регистрации"})
		return
	}

	token, role, err := h.authSvc.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Регистрация успешна, но вход не удался"})
		return
	}

	h.setAuthCookie(c, token)

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Регистрация успешна",
		"role":     role,
	})
	
	c.Redirect(http.StatusFound, "/")
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	token, role, err := h.authSvc.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный логин или пароль"})
		return
	}

	h.setAuthCookie(c, token)

	c.JSON(http.StatusOK, gin.H{
		"message":  "Вход выполнен",
		"role":     role,
	})
	
	c.Redirect(http.StatusFound, "/profile")
}

func (h *AuthHandler) Logout(c *gin.Context) {
	c.SetCookie(
		"FRAMEBY_ACCESS_TOKEN",
		"",
		time.Now().AddDate(0, 0, -1).Second(),
		"/",
		"localhost",
		false,
		false,
	)

	c.JSON(http.StatusOK, gin.H{"message": "Выход выполнен"})
	c.Redirect(http.StatusFound, "/login")
}

func (h *AuthHandler) CheckAuth(c *gin.Context) {
	userID, _ := c.Get("userID")
	role, _ := c.Get("role")

	c.JSON(http.StatusOK, gin.H{
		"authorized": true,
		"userID":     userID,
		"role":       role,
	})
}

// Вспомогательный метод для установки куки, чтобы не дублировать код
func (h *AuthHandler) setAuthCookie(c *gin.Context, token string) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		"FRAMEBY_ACCESS_TOKEN",
		token,
		time.Now().AddDate(0, 1, 0).Second(),
		"/",
		"localhost",
		false,
		false, // HttpOnly
	)
}