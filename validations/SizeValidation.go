package validations

import (
	"errors"
)

// Validates text size of name and brand than can't be less than 3 and greater than 50
func SizeValidation(s string) error {

	if len(s) < 3 || len(s) > 50 {
		err := errors.New(("The size of the name or brand cannot be less than 3 or greater than 50"))
		return err
	}

	return nil
}
