package apperrors

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Field is used to represent a validation error for a specific field
type Field struct {
	Field string `json:"field"`
	Err   string `json:"error"`
}

// FieldErrors is a collection of FieldError
type FieldErrors []Field

// NewFieldErrors returns a FieldErrors with a single FieldError
func NewFieldErrors(field string, err error) FieldErrors {
	return FieldErrors{
		{
			Field: field,
			Err:   err.Error(),
		},
	}
}

// Error returns the string representation of the FieldErrors
func (fe FieldErrors) Error() string {
	b, err := json.Marshal(fe)
	if err != nil {
		return fmt.Sprintf("could not parse FieldErrors: %v", err)
	}
	return string(b)
}

// ToFieldErrors returns the underlying error of a FieldErrors
// returns true if the error is a FieldErrors
func ToFieldErrors(err error) (FieldErrors, bool) {
	var fe FieldErrors
	if errors.As(err, &fe) {
		return fe, true
	}
	return nil, false
}
