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
	return fmt.Sprintf("NotFoundError: %s not found", e.Resource)
}

type ValidationError struct {
	Field   string
	Message string
}

type DuplicateEntryError struct {
	Resource string
	Field    string
	Value    string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("ValidationError: %s - %s", e.Field, e.Message)
}

func NewInternalError(message string) error {
	return &InternalError{Message: message}
}

func NewNotFoundError(resource string) error {
	return &NotFoundError{Resource: resource}
}

func NewValidationError(field, message string) error {
	return &ValidationError{Field: field, Message: message}
}

func (e *DuplicateEntryError) Error() string {
	return fmt.Sprintf("%s with %s '%s' already exists", e.Resource, e.Field, e.Value)
}

func NewDuplicateEntryError(resource, field, value string) error {
	return &DuplicateEntryError{
		Resource: resource,
		Field:    field,
		Value:    value,
	}
}
