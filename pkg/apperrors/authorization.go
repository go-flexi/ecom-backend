package apperrors

import "errors"

// Authorization is used to represent an authorization error
type Authorization struct {
	reason string
}

// NewAuthorization returns a new Authorization error
func NewAuthorization(reason string) error {
	return Authorization{reason: reason}
}

// Error returns the string representation of the Authorization error
func (ae Authorization) Error() string {
	return ae.reason
}

// IsAuthorization checks if an error is a Authorization error
func IsAuthorization(err error) bool {
	var a Authorization
	return errors.As(err, &a)
}
