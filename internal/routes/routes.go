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

type Routes struct {
	Engine      *gin.Engine
	RedisClient *redis.Client
	Handlers    *wirex.Handlers
}

func (a *Routes) RegisterRouters() error {
	// Load configuration
	cfg := config.LoadConfigs()

	gAPI := a.Engine.Group(ApiPrefix)

	gAPI.GET("/health", func(c *gin.Context) {
		ctx := context.Background()

		// Periksa Redis
		if err := a.RedisClient.Ping(ctx).Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"redis": "unavailable"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"redis":  "connected",
		})
	})

	// Login
	gAPI.POST("/login", a.Handlers.LoginApi.Login)

	// Using JWT to access
	gAPI.Use(middleware.AuthMiddleware(a.RedisClient, cfg.JWTConfig.JWTSecretKey))
	gAPI.POST("/logout", a.Handlers.LoginApi.Logout)

	user := gAPI.Group("users")
	{
		user.GET("", a.Handlers.UserApi.Query)
		user.GET(":id", a.Handlers.UserApi.Get)
		user.POST("", a.Handlers.UserApi.Create)
		user.PUT(":id", middleware.AuthorizationMiddleware("id"), a.Handlers.UserApi.Update)
		user.PUT(":id/reset-pwd", middleware.AuthorizationMiddleware("id"), a.Handlers.LoginApi.UpdatePassword)
		user.DELETE(":id", middleware.AuthorizationMiddleware("id"), a.Handlers.UserApi.Delete)
	}

	return nil
}
