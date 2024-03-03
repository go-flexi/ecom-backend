package validate

import (
	"errors"
	"fmt"
	"testing"
)

func TestIsFieldErrors(t *testing.T) {
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
			actual := IsFieldErrors(tc.err)
			if actual != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, actual)
			}
		})
	}
}
