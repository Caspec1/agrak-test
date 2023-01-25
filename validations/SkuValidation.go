package validations

import (
	"errors"
	"strconv"
)

// Convert string to int for validate the sku can't be less than 1000000 and greater than 99999999
func SkuValidation(s string) (string, error) {

	numberSku, _ := strconv.Atoi(s)

	if numberSku < 1000000 || numberSku > 99999999 {
		err := errors.New(("Sku can't be less than 1000000 and greater than 99999999"))
		return "", err
	}

	falSku := strconv.Itoa(numberSku)

	return "FAL-" + falSku, nil
}
