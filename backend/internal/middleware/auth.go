package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lanxre/frameby/internal/services"
)

type AuthMiddleware struct {
	tokenSvc *services.TokenService
}

func NewAuthMiddleware(tokenSvc *services.TokenService) *AuthMiddleware {
	return &AuthMiddleware{tokenSvc: tokenSvc}
}

func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("FRAMEBY_ACCESS_TOKEN")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authentication cookie missing"})
			return
		}

		claims, err := m.tokenSvc.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		c.Set(UserIDKey, claims.UserID)
		c.Set(UserEmailKey, claims.Email)
		c.Set(UserRoleKey, claims.Role)
		c.Next()
	}
}

func (m *AuthMiddleware) RequireRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get(UserRoleKey)
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "role not found"})
			return
		}

		if role != requiredRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
			return
		}

		c.Next()
	}
}
