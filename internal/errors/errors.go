package errors

import "fmt"

type InternalError struct {
	Message string
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("InternalError: %s", e.Message)
}

type NotFoundError struct {
	Resource string
	ID       int
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("NotFoundError: %s with ID %d not found", e.Resource, e.ID)
}

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("ValidationError: %s - %s", e.Field, e.Message)
}

func NewInternalError(message string) error {
	return &InternalError{Message: message}
}

func NewNotFoundError(resource string, id int) error {
	return &NotFoundError{Resource: resource, ID: id}
}

func NewValidationError(field, message string) error {
	return &ValidationError{Field: field, Message: message}
}
