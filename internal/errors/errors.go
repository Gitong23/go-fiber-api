package errors

import "github.com/gofiber/fiber/v2"

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Code:    fiber.StatusUnauthorized,
		Message: message,
	}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{
		Code:    fiber.StatusBadRequest,
		Message: message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    fiber.StatusNotFound,
		Message: message,
	}
}
