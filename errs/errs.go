package errs

import (
	"github.com/gofiber/fiber/v2"
)

type AppError struct {
	Code        int
	Message     string
	JSONMessage map[string]interface{}
}

func (e AppError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) error {
	return AppError{
		Code:    fiber.StatusNotFound,
		Message: message,
		JSONMessage: fiber.Map{
			"message": message,
		},
	}
}

func NewUnexpectedError() error {
	return AppError{
		Code:    fiber.StatusNotFound,
		Message: "Unexpected error",
		JSONMessage: fiber.Map{
			"message": "Unexpected error",
		},
	}
}

// func NewValidationError(message string) error {
// 	return AppError{
// 		Code:    http.StatusUnprocessableEntity,
// 		Message: message,
// 	}
// }
