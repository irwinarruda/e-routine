package app_error

import "fmt"

type statusCode int

const (
	NoStatus            statusCode = -1
	InternalServerError statusCode = 500
	BadRequest          statusCode = 404
)

type AppError struct {
	message    string
	statusCode statusCode
}

func (this AppError) Error() string {
	return fmt.Sprintf("AppError: %v", this.message)
}

func New(message string, statusCode statusCode) error {
	return AppError{message, statusCode}
}
