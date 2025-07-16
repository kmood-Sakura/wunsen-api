// internal/core/utils/validation.go
package utils

import "go-api/internal/infra/log"

func ValidateGender(gender string) bool {
	isValid := gender == "male" || gender == "female"
	if !isValid {
		log.APILogger.Printf("[VALIDATION-ERROR] Invalid gender provided: %s", gender)
	}
	return isValid
}