package validations

import (
	"errors"
	"testing"
)

// Test Size Validation
func TestSizeValidation(t *testing.T) {
	testCases := []struct {
		Name          string
		s             string
		ExpectedError error
	}{
		{
			Name:          "Correct size",
			s:             "500 Zapatilla Urbana Mujer",
			ExpectedError: nil,
		},
		{
			Name:          "Less than 3",
			s:             "Al",
			ExpectedError: errors.New(("The size of the name or brand cannot be less than 3 or greater than 50")),
		},
		{
			Name:          "Greater than 50",
			s:             "This is a text with more than 50 letters for the testing",
			ExpectedError: errors.New(("The size of the name or brand cannot be less than 3 or greater than 50")),
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			err := SizeValidation(tc.s)

			if len(tc.s) < 3 || len(tc.s) > 50 {
				t.Logf("Expected %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
