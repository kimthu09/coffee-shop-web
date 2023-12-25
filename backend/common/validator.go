package common

import (
	"net/mail"
	"regexp"
)

func ValidateEmptyString(s string) bool {
	return s == ""
}

func ValidatePhone(s string) bool {
	if len(s) != 10 && len(s) != 11 {
		return false
	}

	pattern := `\d`

	regex := regexp.MustCompile(pattern)

	return regex.MatchString(s)
}

func ValidateEmail(s string) bool {
	_, err := mail.ParseAddress(s)
	return err == nil
}

func ValidateNotNilId(id *string) bool {
	if id == nil || len(*id) == 0 || len(*id) > MaxLengthIdCanGenerate {
		return false
	}
	return true
}

func ValidateId(id *string) bool {
	if id == nil || len(*id) == 0 {
		return true
	}
	if len(*id) > MaxLengthIdCanGenerate {
		return false
	}
	return true
}

func ValidateNegativeNumber(number interface{}) bool {
	switch v := number.(type) {
	case int:
		return v < 0
	case int8:
		return v < 0
	case int16:
		return v < 0
	case int64:
		return v < 0
	case float32:
		return v < 0
	case float64:
		return v < 0
	default:
		return false
	}
}

func ValidateNotPositiveNumber(number interface{}) bool {
	switch v := number.(type) {
	case int:
		return v <= 0
	case int8:
		return v <= 0
	case int16:
		return v <= 0
	case int64:
		return v <= 0
	case float32:
		return v <= 0
	case float64:
		return v <= 0
	default:
		return false
	}
}

func ValidatePassword(pass *string) bool {
	return pass != nil && len(*pass) >= 6
}

func ValidatePositiveNumber(number interface{}) bool {
	switch v := number.(type) {
	case int:
		return v > 0
	case int8:
		return v > 0
	case int16:
		return v > 0
	case int64:
		return v > 0
	case float32:
		return v > 0
	case float64:
		return v > 0
	default:
		return false
	}
}
