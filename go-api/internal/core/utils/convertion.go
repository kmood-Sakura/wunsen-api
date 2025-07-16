// internal/core/utils/conversion.go
package utils

func ConvertCmToM(cm float64) float64 {
	meters := cm / 100.0
	return meters
}