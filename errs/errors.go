package errs

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty" xml:"code"`
	Message string `json:"message" xml:"message"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func NewNotFoundError(m string) *AppError {
	return &AppError{
		Message: m,
		Code:    http.StatusNotFound,
	}
}

func NewInternalSystemError(m string) *AppError {
	return &AppError{
		Message: m,
		Code:    http.StatusInternalServerError,
	}
}

func NewValidationError(m string) *AppError {
	return &AppError{
		Message: m,
		Code:    http.StatusUnprocessableEntity,
	}
}
