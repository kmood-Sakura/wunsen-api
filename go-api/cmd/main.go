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

	// Setup routes
	router.SetupRoutes(r)

	// Start server
	addr := cfg.Application.Host + ":" + cfg.Application.Port
	log.Printf("Server starting on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}