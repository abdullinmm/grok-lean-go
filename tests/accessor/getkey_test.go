package accessor_test

import (
	"errors"
	"testing"

	"github.com/abdullinmm/grok-lean-go/internal/accessor"
)

func TestGetKey(t *testing.T) {

	// Created and filling the card with test data
	m := make(map[string]int)
	m["Marsel"] = 44

	test := []struct {
		name     string
		key      string
		value    int
		expected error
	}{
		{
			name:     "valid map key",
			key:      "Marsel",
			value:    44,
			expected: nil},
		{
			name:     "invalid map key",
			key:      "Mrl",
			expected: accessor.ErrInvalidMapKey},
	}

	// We pass each test
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			_, err := accessor.GetKey(m, tt.key)
			if !errors.Is(err, tt.expected) {
				t.Errorf("expected key %q to exist in the map", tt.key)
			}
		})
	}
}
