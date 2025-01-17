package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Example routes
	router.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	return router
}
