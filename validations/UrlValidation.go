package validations

import (
	"errors"
	"strings"
)

// Check an url direction
func UrlValidation(s string) error {

	checkHttp := strings.HasPrefix(s, "http")

	if !checkHttp {
		err := errors.New(("Url must starts with http"))
		return err
	}

	return nil
}
