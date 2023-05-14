package util

import (
	"dbo-test/config"
	"fmt"
	"regexp"
)

func ValidateEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("%s: wrong format email", config.Configure.Source.Name)
	}
	return nil
}
