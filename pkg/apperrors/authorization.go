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

// ToAuthorization returns the underlying error of a Authorization error
// returns true if the error is a Authorization error
func ToAuthorization(err error) (error, bool) {
	var a Authorization
	if errors.As(err, &a) {
		return a, true
	}
	return err, false
}
