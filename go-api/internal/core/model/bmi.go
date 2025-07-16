// internal/core/model/bmi.go
package model

type BMIRequest struct {
	Weight float64 `json:"weight" binding:"required,gt=0"`
	Height float64 `json:"height" binding:"required,gt=0"`
	Sex    string  `json:"sex" binding:"required,oneof=male female"`
}

type BMIResponse struct {
	BMI       float64 `json:"BMI"`
	BMIStatus string  `json:"BMI_STATUS"`
}

type API2Request struct {
	Weight float64 `json:"weight"`
	Height float64 `json:"height"`
}

type API2Response struct {
	BMI float64 `json:"BMI"`
}