package routes

import (
	"finance-tracker/internal/wirex"
	"finance-tracker/pkg/config"
	"finance-tracker/pkg/middleware"
	"finance-tracker/pkg/util"

	"github.com/gin-gonic/gin"
)

const (
	ApiPrefix = "/api/v1/"
)

type Routes struct{}

func (a *Routes) RegisterRouters(e *gin.Engine, handlers *wirex.Handlers) error {
	// Load configuration
	cfg := config.LoadConfigs()

	gAPI := e.Group(ApiPrefix)

	gAPI.GET("/health", func(c *gin.Context) {
		util.ResOK(c)
	})

	// Login Logout
	gAPI.POST("/login", handlers.LoginApi.Login)
	gAPI.POST("/logout", handlers.LoginApi.Logout)

	// Using JWT to access
	gAPI.Use(middleware.AuthMiddleware(cfg.JWTConfig.JWTSecretKey)) // Ganti dengan secret JWT nyata
	user := gAPI.Group("users")
	{
		user.GET("", handlers.UserApi.Query)
		user.GET(":id", handlers.UserApi.Get)
		user.POST("", handlers.UserApi.Create)
		user.PUT(":id", handlers.UserApi.Update)
		user.DELETE(":id", handlers.UserApi.Delete)
	}

	return nil
}
