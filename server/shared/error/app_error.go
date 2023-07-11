package error

import "fmt"

type AppError struct {
	message    string
	statusCode int
}

func (this AppError) Error() string {
	return fmt.Sprintf("AppError: %v", this.message)
}

func New(message string, statusCode int) error {
	return AppError{message, statusCode}
}
