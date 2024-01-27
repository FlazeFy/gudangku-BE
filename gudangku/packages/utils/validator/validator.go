package validator

import (
	"strings"
	"time"
)

func GetValidateEmail(val string) bool {
	return strings.HasSuffix(val, "@gmail.com")
}

func GetValidationLength(col string) (int, int) {
	if col == "username" {
		return 6, 36
	} else if col == "email" {
		return 10, 75
	} else if col == "password" {
		return 6, 36
	} else if col == "first_name" {
		return 1, 36
	} else if col == "last_name" {
		return 0, 36
	} else if col == "valid_until" {
		yearNow := time.Now().Year()
		max := yearNow + 6
		min := yearNow - 6
		return min, max
	}
	return 0, 0
}
