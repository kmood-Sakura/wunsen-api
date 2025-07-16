// internal/handler/bmi_handler.go
package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go-api/internal/core/model"
	"go-api/internal/core/service"
	"go-api/internal/infra/config"
	"go-api/internal/infra/log"
)

type BMIHandler struct {
	bmiService *service.BMIService
}

func NewBMIHandler(cfg *config.Config) *BMIHandler {
	return &BMIHandler{
		bmiService: service.NewBMIService(cfg),
	}
}

func (h *BMIHandler) CalculateBMI(c *gin.Context) {
	startTime := time.Now()
	clientIP := c.ClientIP()

	// Step 1: Get request body and bind JSON
	var req model.BMIRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		duration := time.Since(startTime)
		
		errorResponse := gin.H{
			"error": "Invalid request format",
			"details": err.Error(),
		}
		log.LogError(clientIP, err.Error())
		log.LogBMIResponse(clientIP, http.StatusBadRequest, errorResponse, duration)
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	log.LogBMIRequest(clientIP, req)

	// Process BMI calculation through service
	response, err := h.bmiService.ProcessBMIRequest(&req)
	duration := time.Since(startTime)

	if err != nil {
		errorResponse := gin.H{
			"error": err.Error(),
		}
		log.LogError(clientIP, err.Error())
		log.LogBMIResponse(clientIP, http.StatusBadRequest, errorResponse, duration)
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Return successful response
	log.LogBMIResponse(clientIP, http.StatusOK, response, duration)
	c.JSON(http.StatusOK, response)
}