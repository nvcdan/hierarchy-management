package errors

type AuthenticationError struct {
	Message string
}

func (e *AuthenticationError) Error() string {
	return e.Message
}

func NewAuthenticationError(message string) error {
	return &AuthenticationError{Message: message}
}
