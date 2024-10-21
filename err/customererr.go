package err

import (
	"fmt"
)

// CustomError represents a custom error type.
type CustomError struct {
	Code    int
	Message string
}

// Error implements the error interface for CustomError.
func (e *CustomError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// New creates a new CustomError.
func New(code int, message string) error {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}
