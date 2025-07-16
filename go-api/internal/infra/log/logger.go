// internal/infra/log/logger.go
package log

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	APILogger *log.Logger
)

func init() {
	InitializeLogger()
}

func InitializeLogger() {
	logDir := "./log"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Printf("Failed to create log directory: %v", err)
		// Fallback to stdout only
		setupLoggers(os.Stdout)
		return
	}

	// Create or open log file
	logFile := filepath.Join(logDir, "api-log.txt")
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Failed to open log file: %v", err)
		// Fallback to stdout only
		setupLoggers(os.Stdout)
		return
	}

	// Setup loggers with both file and stdout output
	multiWriter := io.MultiWriter(os.Stdout, file)
	setupLoggers(multiWriter)
}

func setupLoggers(writer io.Writer) {
	APILogger = log.New(writer, "", log.Ldate|log.Ltime|log.Lmicroseconds)
}

// LogBMIRequest logs incoming BMI calculation requests
func LogBMIRequest(clientIP string, requestBody interface{}) {
	APILogger.Printf("[REQUEST] BMI calculation from %s | Request: %+v", clientIP, requestBody)
}

// LogBMIResponse logs BMI calculation responses
func LogBMIResponse(clientIP string, statusCode int, responseBody interface{}, duration time.Duration) {
	APILogger.Printf("[RESPONSE] BMI result to %s | Status: %d | Duration: %v | Response: %+v", 
		clientIP, statusCode, duration, responseBody)
}

// LogAPI2Request logs API(2) requests
func LogAPI2Request(url string, requestBody interface{}) {
	APILogger.Printf("[API2-REQUEST] Calling %s | Request: %+v", url, requestBody)
}

// LogAPI2Response logs API(2) responses
func LogAPI2Response(url string, statusCode int, responseBody interface{}, duration time.Duration) {
	APILogger.Printf("[API2-RESPONSE] Response from %s | Status: %d | Duration: %v | Response: %+v", 
		url, statusCode, duration, responseBody)
}

// LogError logs errors during BMI calculation process
func LogError(clientIP string, errorMsg string) {
	APILogger.Printf("[ERROR] BMI calculation failed for %s | Error: %s", clientIP, errorMsg)
}