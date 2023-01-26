package validations

import (
	"errors"
	"testing"
)

// Test Price Validation
func TestPriceValidation(t *testing.T) {
	testCases := []struct {
		Name          string
		p             float64
		ExpectedError error
	}{
		{
			Name:          "Correct size",
			p:             54000.12,
			ExpectedError: nil,
		},
		{
			Name:          "Less than 1",
			p:             0.30,
			ExpectedError: errors.New(("Price can't be less than 1 and greater than 99999999")),
		},
		{
			Name:          "Greater than 99999999",
			p:             999999999.02,
			ExpectedError: errors.New(("Price can't be less than 1 and greater than 99999999")),
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			err := PriceValidation(tc.p)

			if tc.p < 1.00 || tc.p > 99999999.00 {
				t.Logf("Expected %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
