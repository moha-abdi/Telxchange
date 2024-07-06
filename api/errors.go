package api

import "fmt"

type APIError struct {
	Code    string
	Message string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("Error (%s): %s", e.Code, e.Message)
}

func NewAPIError(code string, message string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
	}
}
