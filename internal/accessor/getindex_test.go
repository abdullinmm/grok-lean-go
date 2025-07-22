package accessor_test

import (
	"errors"
	"testing"

	"github.com/abdullinmm/grok-lean-go/internal/accessor"
)

func TestGetIndex(t *testing.T) {

	test := []struct {
		name     string
		index    int
		slice    []int
		expected error
	}{
		{
			name:     "index out of range",
			index:    100,
			slice:    []int{1, 2, 3},
			expected: accessor.ErrIndexOutOfBounds},
		{
			name:     "index is range",
			index:    1,
			slice:    []int{1, 2, 3},
			expected: nil},
		{
			name:     "negative index value",
			index:    -1,
			slice:    []int{1, 2, 3},
			expected: accessor.ErrIndexOutOfBounds},
	}

	// We pass each test
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			value, err := accessor.GetIndex(tt.slice, tt.index)
			if !errors.Is(err, tt.expected) {
				t.Errorf("Expected %s, got %s", tt.expected, err)
			} else if err == nil && tt.expected == nil {
				// Check the value if there is no error
				expectedVal := tt.index
				if value != expectedVal {
					t.Errorf("Expected value %v, got %v", expectedVal, value)
				}
			}
		})
	}
}
