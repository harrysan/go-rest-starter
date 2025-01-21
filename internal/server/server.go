package server

import (
	"log"
	"net/http"
	"time"

	_ "finance-tracker/docs"
	"finance-tracker/internal/mods"
	"finance-tracker/internal/wirex"
	"finance-tracker/pkg/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewServer(handlers *wirex.Handlers, cfg config.App) *http.Server {
	r := gin.Default()

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Setup routes
	mods := mods.Mods{}
	mods.RegisterRouters(r, handlers)

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
