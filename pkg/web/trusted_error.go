package web

import "errors"

// TrustedError is used to represent an error that is trusted and should be returned to the client
type TrustedError struct {
	Err    error
	Status int
}

// NewTrustedError returns a new TrustedError
func NewTrustedError(err error, status int) error {
	return TrustedError{
		Err:    err,
		Status: status,
	}
}

// Error returns the string representation of the TrustedError
func (te TrustedError) Error() string {
	return te.Err.Error()
}

// IsTrustedError checks if an error is a TrustedError
func IsTrustedError(err error) bool {
	var te TrustedError
	return errors.As(err, &te)
}
