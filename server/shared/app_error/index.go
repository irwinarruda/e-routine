package app_error

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type statusCode int

const (
	NoStatus            statusCode = -1
	InternalServerError statusCode = 500
	BadRequest          statusCode = 404
)

type AppError struct {
	message         string
	statusCode      statusCode
	isPublicMessage bool
}

func (this AppError) Error() string {
	return fmt.Sprintf("AppError: %v", this.message)
}

func New(message string, statusCode statusCode) error {
	return AppError{message: message, statusCode: statusCode, isPublicMessage: true}
}

func NewPrivate(message string, statusCode statusCode) error {
	return AppError{message: message, statusCode: statusCode, isPublicMessage: false}
}

func HandleHttp(ctx *fiber.Ctx, err error) error {
	if appError, ok := err.(AppError); ok {
		return ctx.JSON(appError)
	}

	return ctx.JSON(AppError{message: err.Error(), statusCode: InternalServerError})
}
