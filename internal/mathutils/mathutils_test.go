package mathutils

import (
	"testing"
)

func TestMax(t *testing.T) {
	test := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "first value is not greater than the second",
			a:        10,
			b:        5,
			expected: 10,
		},
		{
			name:     "second value is not greater than the first",
			a:        5,
			b:        10,
			expected: 10,
		},
		{
			name:     "both values are equal",
			a:        10,
			b:        10,
			expected: 10,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if Max(tt.a, tt.b) != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, Max(tt.a, tt.b))
			}
		})
	}
}
