package web

import "errors"

type TrustedError struct {
	Err    error
	Status int
}

func NewTrustedError(err error, status int) error {
	return TrustedError{
		Err:    err,
		Status: status,
	}
}

func (te TrustedError) Error() string {
	return te.Err.Error()
}

func IsTrustedError(err error) bool {
	var te TrustedError
	return errors.As(err, &te)
}
