package utils

import "fmt"

type CustomError struct {
	ErrorMessage string
}

func (customerError *CustomError) Error() string {
	return fmt.Sprint(customerError.ErrorMessage)
}
