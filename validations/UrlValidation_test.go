package validations

import (
	"errors"
	"strings"
	"testing"
)

func TestUrlValidation(t *testing.T) {
	testCases := []struct {
		Name          string
		s             string
		ExpectedError error
	}{
		{
			Name:          "Correct size",
			s:             "https://localhost:3000",
			ExpectedError: nil,
		},
		{
			Name:          "Less than 1",
			s:             "localhost:3000",
			ExpectedError: errors.New(("Url must starts with http")),
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			err := UrlValidation(tc.s)

			checkHttp := strings.HasPrefix(tc.s, "http")

			if !checkHttp {
				t.Logf("Expected %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
