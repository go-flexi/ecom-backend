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

// ToTrustedError returns the underlying error of a TrustedError
// returns true if the error is a TrustedError
func ToTrustedError(err error) (TrustedError, bool) {
	var te TrustedError
	if errors.As(err, &te) {
		return te, true
	}
	return TrustedError{}, false
}
