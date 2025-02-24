package server

import (
	"log"
	"net/http"
	"time"

	_ "finance-tracker/docs"
	"finance-tracker/internal/routes"
	"finance-tracker/internal/wirex"
	"finance-tracker/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewServer(handlers *wirex.Handlers, redisClient *redis.Client, cfg config.App) *http.Server {
	r := gin.Default()

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes := routes.Routes{
		Engine:      r,
		RedisClient: redisClient,
		Handlers:    handlers,
	}
	routes.RegisterRouters()

	// Configure server
	server := &http.Server{
		Addr:         cfg.AppConfig.Port,
		Handler:      r,
		ReadTimeout:  time.Duration(cfg.AppTimeoutConfig.Read * int(time.Second)),
		WriteTimeout: time.Duration(cfg.AppTimeoutConfig.Write * int(time.Second)),
		IdleTimeout:  time.Duration(cfg.AppTimeoutConfig.Idle * int(time.Second)),
	}

	return server
}

func StartServer(server *http.Server) {
	log.Println("Starting server on port " + server.Addr + "...")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not start server: %v", err)
	}
}
