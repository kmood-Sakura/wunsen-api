// internal/core/service/bmi_service.go
package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"go-api/internal/core/model"
	"go-api/internal/core/utils"
	"go-api/internal/infra/config"
	"go-api/internal/infra/log"
)

type BMIService struct {
	config *config.Config
	client *http.Client
}

func NewBMIService(cfg *config.Config) *BMIService {
	return &BMIService{
		config: cfg,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *BMIService) ProcessBMIRequest(req *model.BMIRequest) (*model.BMIResponse, error) {
	// Step 1: Validate sex/gender
	if !utils.ValidateGender(req.Sex) {
		return nil, errors.New("invalid gender: must be 'male' or 'female'")
	}

	// Step 2: Convert height from cm to m
	heightInMeters := utils.ConvertCmToM(req.Height)

	// Step 3: Call API(2) to calculate BMI magnitude
	api2Req := &model.API2Request{
		Weight: req.Weight,
		Height: heightInMeters,
	}

	api2Resp, err := s.callAPI2(api2Req)
	if err != nil {
		return nil, fmt.Errorf("failed to call API(2): %v", err)
	}

	// Step 4: Compare BMI with criteria to get status (gender-specific)
	bmiStatus := utils.GetBMIStatus(api2Resp.BMI, req.Sex)

	// Step 5: Prepare response
	response := &model.BMIResponse{
		BMI:       api2Resp.BMI,
		BMIStatus: bmiStatus,
	}

	return response, nil
}

func (s *BMIService) callAPI2(req *model.API2Request) (*model.API2Response, error) {
	startTime := time.Now()

	// Prepare request body
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	// Build API(2) URL
	api2URL := fmt.Sprintf("http://%s:%s/%s", 
		s.config.API2.Host, s.config.API2.Port, s.config.API2.Endpoint)

	log.LogAPI2Request(api2URL, req)

	// Create HTTP request
	httpReq, err := http.NewRequest("POST", api2URL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	// Make the request
	resp, err := s.client.Do(httpReq)
	duration := time.Since(startTime)

	if err != nil {
		return nil, fmt.Errorf("failed to call API(2): %v", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		log.LogAPI2Response(api2URL, resp.StatusCode, nil, duration)
		return nil, fmt.Errorf("API(2) returned status %d", resp.StatusCode)
	}

	// Parse response
	var api2Resp model.API2Response
	if err := json.NewDecoder(resp.Body).Decode(&api2Resp); err != nil {
		return nil, fmt.Errorf("failed to decode API(2) response: %v", err)
	}

	log.LogAPI2Response(api2URL, resp.StatusCode, api2Resp, duration)

	return &api2Resp, nil
}