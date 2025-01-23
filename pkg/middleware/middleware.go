package middleware

import (
	"context"
	"net/http"
	"strings"

	"finance-tracker/pkg/jwt"
	rds "finance-tracker/pkg/redis"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// AuthMiddleware godoc
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func AuthMiddleware(redisClient *redis.Client, secretKey string) gin.HandlerFunc {
	jwtManager := jwt.NewJWTManager(secretKey, 0) // Expiry hanya untuk generate token
	return func(c *gin.Context) {
		// Health check Redis sebelum melanjutkan
		if err := redisClient.Ping(context.Background()).Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Authmiddleware, Redis unavailable"})
			c.Abort()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Invalid token format."})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Cek token di Redis blacklist
		isBlacklisted, err := rds.IsTokenBlacklisted(redisClient, token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking token"})
			c.Abort()
			return
		}

		if isBlacklisted {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
			c.Abort()
			return
		}

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
