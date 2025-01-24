package middleware

import (
	"context"
	"net/http"
	"strconv"
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

// AuthorizationMiddleware memastikan hanya user yang login dapat memproses request tertentu
func AuthorizationMiddleware(paramKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil user_id dari JWT (dari context)
		userIDFromToken, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Ambil user_id dari parameter (dari URL, query, atau lainnya)
		paramValue := c.Param(paramKey) // Param diambil dari URL
		userIDFromParam, err := strconv.Atoi(paramValue)
		if err != nil || userIDFromParam <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			c.Abort()
			return
		}

		// Cocokkan user_id dari token dan URL
		if userIDFromToken != uint(userIDFromParam) {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to access this resource"})
			c.Abort()
			return
		}

		c.Next() // Lanjutkan ke handler berikutnya jika validasi berhasil
	}
}
