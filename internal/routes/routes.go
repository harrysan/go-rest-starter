package routes

import (
	"context"
	"finance-tracker/internal/wirex"
	"finance-tracker/pkg/config"
	"finance-tracker/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

const (
	ApiPrefix = "/api/v1/"
)

type Routes struct{}

func (a *Routes) RegisterRouters(e *gin.Engine, redisClient *redis.Client, handlers *wirex.Handlers) error {
	// Load configuration
	cfg := config.LoadConfigs()

	gAPI := e.Group(ApiPrefix)

	gAPI.GET("/health", func(c *gin.Context) {
		ctx := context.Background()

		// Periksa Redis
		if err := redisClient.Ping(ctx).Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"redis": "unavailable"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"redis":  "connected",
		})
	})

	// Login
	gAPI.POST("/login", handlers.LoginApi.Login)

	// Using JWT to access
	gAPI.Use(middleware.AuthMiddleware(redisClient, cfg.JWTConfig.JWTSecretKey))
	gAPI.POST("/logout", handlers.LoginApi.Logout)

	user := gAPI.Group("users")
	{
		user.GET("", handlers.UserApi.Query)
		user.GET(":id", handlers.UserApi.Get)
		user.POST("", handlers.UserApi.Create)
		user.PUT(":id", middleware.AuthorizationMiddleware("id"), handlers.UserApi.Update)
		user.PUT(":id/reset-pwd", middleware.AuthorizationMiddleware("id"), handlers.LoginApi.UpdatePassword)
		user.DELETE(":id", middleware.AuthorizationMiddleware("id"), handlers.UserApi.Delete)
	}

	return nil
}
