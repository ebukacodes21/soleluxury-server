package validate

import "fmt"

func ValidateString(value string, min, max int) error {
	n := len(value)
	if n < min || n > max {
		return fmt.Errorf("must contain from %d-%d characters", min, max)
	}
	return nil
}

func ValidateStoreName(value string) error {
	return ValidateString(value, 3, 100)
}
