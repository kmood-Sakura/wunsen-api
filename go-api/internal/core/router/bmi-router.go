// internal/core/router/router.go
package router

import (
	"github.com/gin-gonic/gin"
	"go-api/internal/handler"
	"go-api/internal/infra/config"
	"go-api/internal/infra/log"
)

func SetupRoutes(r *gin.Engine) {
	// Load config for handlers
	cfg := config.LoadConfig()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		clientIP := c.ClientIP()
		response := gin.H{
			"status":  "healthy",
			"message": "api is healthy",
		}
		
		log.APILogger.Printf("[HEALTH-RESPONSE] Health check to %s | Status: 200 | Response: %+v", clientIP, response)
		c.JSON(200, response)
	})

	// BMI API endpoints
	bmiHandler := handler.NewBMIHandler(cfg)
	
	api := r.Group("/api")
	{
		api.POST("/bmi", bmiHandler.CalculateBMI)
	}
}