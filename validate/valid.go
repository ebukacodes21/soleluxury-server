package validate

import (
	"errors"
	"fmt"
	"math/rand"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
)

const alpha = "abcdefghijklmnopqrstuvwxyz"

var (
	isUsernameValid = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
)

func ValidateString(value string, min, max int) error {
	n := len(value)
	if n < min || n > max {
		return fmt.Errorf("must contain from %d-%d characters", min, max)
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isUsernameValid(value) {
		return fmt.Errorf("must contain only letters, digits or underscore")
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 8, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}

	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("not a valid email address")
	}
	return nil
}

func ValidateStoreName(value string) error {
	return ValidateString(value, 3, 100)
}

func RandomString(n int) string {
	var sb strings.Builder

	for i := 0; i < n; i++ {
		char := alpha[rand.Intn(len(alpha))]
		sb.WriteByte(char)
	}

	return sb.String()
}

func ValidateId(value int64) error {
	if value <= 0 {
		return fmt.Errorf("value must be a positive integer")
	}
	return nil
}

func ValidateUrl(value string) error {
	_, err := url.ParseRequestURI(value)
	if err != nil {
		return errors.New("invalid URL")
	}
	return nil
}
