package validations

import (
	"errors"
)

// Check price can't be less than 1 and greater than 99999999
func PrinceValidation(p float64) error {

	if p < 1.00 || p > 99999999.00 {
		err := errors.New(("Price can't be less than 1 and greater than 99999999"))
		return err
	}

	return nil
}
