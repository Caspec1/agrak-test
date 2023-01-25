package validations

import (
	"errors"
	"strings"

	"github.com/lib/pq"
)

func OtherUrlValidation(a pq.StringArray) error {

	for _, url := range a {
		checkHttp := strings.HasPrefix(url, "http")
		if !checkHttp {
			err := errors.New(("Url must starts with http"))
			return err
		}
	}

	return nil
}
