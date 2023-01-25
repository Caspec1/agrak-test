package validations

import (
	"errors"
	"testing"
)

func TestSkuValidation(t *testing.T) {

	testCases := []struct {
		Name          string
		s             string
		Expected      string
		ExpectedError error
	}{
		{
			Name:          "Correct SKU",
			s:             "12345678",
			Expected:      "FAL-12345678",
			ExpectedError: nil,
		},
		{
			Name:          "Incorrect SKU",
			s:             "123456789",
			Expected:      "",
			ExpectedError: errors.New("Sku can't be less than 1000000 and greater than 99999999"),
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			result, err := SkuValidation(tc.s)

			if result != tc.Expected && err != nil {
				t.Error("the conversion is incorrect, got: " + result + ", want: " + tc.Expected)
			}
		})
	}
}
