// cmd/main.go
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-api/internal/core/router"
	"go-api/internal/infra/config"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize Gin router
	r := gin.Default()

	// Setup CORS middleware to allow requests from any origin
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

	// Setup routes
	router.SetupRoutes(r)

	// Determine host based on configuration
	var host string
	if cfg.Application.Host == "any" {
		host = "0.0.0.0" // Bind to all interfaces
		log.Printf("Server starting on %s:%s (accessible from any IP)", host, cfg.Application.Port)
	} else {
		host = cfg.Application.Host // Use specific IP from config
		log.Printf("Server starting on %s:%s (accessible from specific IP only)", host, cfg.Application.Port)
	}

	// Start server
	addr := host + ":" + cfg.Application.Port
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}