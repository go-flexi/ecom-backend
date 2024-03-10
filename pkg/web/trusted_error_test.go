package web

import (
	"fmt"
	"testing"
)

func TestToTrustedError(t *testing.T) {
	testTable := map[string]struct {
		err    error
		result bool
	}{
		"trusted error": {
			err:    NewTrustedError(nil, 0),
			result: true,
		},
		"trusted error with wrapped error": {
			err:    fmt.Errorf("wrapped: %w", NewTrustedError(nil, 0)),
			result: true,
		},
		"not trusted error": {
			err:    nil,
			result: false,
		},
	}

	for name, tt := range testTable {
		t.Run(name, func(t *testing.T) {
			if _, ok := ToTrustedError(tt.err); ok != tt.result {
				t.Errorf("expected %v but got %v", tt.result, ok)
			}
		})
	}
}
