package middleware

import (
	"net/http"
	"strings"

	"finance-tracker/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware godoc
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func AuthMiddleware(secretKey string) gin.HandlerFunc {
	jwtManager := jwt.NewJWTManager(secretKey, 0) // Expiry hanya untuk generate token
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Invalid token format."})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := jwtManager.VerifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Invalid or expired token"})
			c.Abort()
			return
		}

		// Simpan informasi user di context
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)

		c.Next()
	}
}
