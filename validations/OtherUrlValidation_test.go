package validations

import (
	"errors"
	"strings"
	"testing"
)

// Test url validation
func TestOtherUrlValidation(t *testing.T) {
	testCases := []struct {
		Name          string
		s             []string
		ExpectedError error
	}{
		{
			Name:          "Correct size",
			s:             []string{"http://localhost:3000", "http://javiermirandadev.com"},
			ExpectedError: nil,
		},
		{
			Name:          "Less than 1",
			s:             []string{"google.com", "twitter.com"},
			ExpectedError: errors.New(("Url must starts with http")),
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			err := OtherUrlValidation(tc.s)

			for _, url := range tc.s {
				checkHttp := strings.HasPrefix(url, "http")
				if !checkHttp {
					t.Logf("Expected %v, got %v", tc.ExpectedError, err)
				}
			}
		})
	}
}
