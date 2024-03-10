package apperrors

import (
	"errors"
	"fmt"
	"testing"
)

func TestToFieldErrors(t *testing.T) {
	testTable := map[string]struct {
		err      error
		expected bool
	}{
		"FieldErrors": {
			err:      NewFieldErrors("test", errors.New("test")),
			expected: true,
		},
		"NonFieldErrors": {
			err:      errors.New("test"),
			expected: false,
		},
		"Nil": {
			err:      nil,
			expected: false,
		},
		"chain error": {
			err:      fmt.Errorf("test: %w", NewFieldErrors("test", errors.New("test"))),
			expected: true,
		},
	}

	for tn, tc := range testTable {
		t.Run(tn, func(t *testing.T) {
			_, ok := ToFieldErrors(tc.err)
			if ok != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, ok)
			}
		})
	}
}
