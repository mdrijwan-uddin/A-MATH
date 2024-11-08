package err

import (
	"fmt"
)

// AmathError represents a custom error type.
type AmathError struct {
	Code    int
	Message string
}

// Error implements the error interface for CustomError.
func (e *AmathError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// New creates a new CustomError.
func New(code int, message string) error {
	return &AmathError{
		Code:    code,
		Message: message,
	}
}
