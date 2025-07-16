// internal/core/utils/bmi_criteria.go
package utils

func GetBMIStatus(bmi float64, gender string) string {
	var status string
	
	if gender == "female" {
		// Female BMI criteria (slightly different thresholds)
		switch {
		case bmi < 18.5:
			status = "Underweight"
		case bmi >= 18.5 && bmi < 24:
			status = "Normal weight"
		case bmi >= 24 && bmi < 29:
			status = "Overweight"
		default:
			status = "Obese"
		}
	} else {
		// Male BMI criteria (standard WHO criteria)
		switch {
		case bmi < 18.5:
			status = "Underweight"
		case bmi >= 18.5 && bmi < 25:
			status = "Normal weight"
		case bmi >= 25 && bmi < 30:
			status = "Overweight"
		default:
			status = "Obese"
		}
	}
	return status
}