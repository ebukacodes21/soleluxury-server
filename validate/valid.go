package validate

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"net/mail"
	"net/url"
	"regexp"
	"strings"

	"github.com/ebukacodes21/soleluxury-server/pb"
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

func ValidateName(value string) error {
	return ValidateString(value, 3, 100)
}

func ValidateDescription(value string) error {
	return ValidateString(value, 3, 100)
}

func ValidatePrice(value float32) error {
	if math.IsNaN(float64(value)) {
		return fmt.Errorf("value cannot be NaN")
	}
	if math.IsInf(float64(value), 0) {
		return fmt.Errorf("value cannot be infinity")
	}
	return nil
}

func ValidateValue(value string) error {
	return ValidateString(value, 1, 1)
}

func ValidateColorValue(value string) error {
	return ValidateString(value, 1, 50)
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

func ValidateUrls(values []*pb.Item) error {
	if len(values) == 0 {
		return errors.New("no URLs provided")
	}

	for _, value := range values {
		if value == nil || value.Url == "" {
			return errors.New("URL is empty")
		}

		_, err := url.ParseRequestURI(value.Url)
		if err != nil {
			return errors.New("invalid URL: " + value.Url)
		}
	}
	return nil
}

func ValidateBool(value bool) error {
	return nil
}
