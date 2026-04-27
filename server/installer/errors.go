package installer

import "errors"

type validationError struct {
	msg string
}

func (e *validationError) Error() string {
	return e.msg
}

func IsValidationError(err error) bool {
	var validationError *validationError
	ok := errors.As(err, &validationError)

	return ok
}

func NewValidationError(msg string) error {
	return &validationError{msg: msg}
}
