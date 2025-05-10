package accessor

import (
	"errors"
	"math/rand/v2"
	"testing"

	"github.com/abdullinmm/grok-lean-go/internal/calculator"
	"github.com/abdullinmm/grok-lean-go/internal/slicecheck"
)

func TestDoubledSlice(t *testing.T) {
	test := []struct {
		name         string
		slice        []int
		doubledSlice []int
		expected     error
	}{
		{
			name:         "Valid test",
			slice:        []int{1, 2, 3, 4, 5},
			doubledSlice: []int{2, 4, 6, 8, 10},
			expected:     nil,
		},
		{
			name:     "Nil slice",
			expected: slicecheck.ErrSliceNil,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			value, err := calculator.DoubleSlice(tt.slice)
			if !errors.Is(err, tt.expected) {
				t.Errorf("Expected %s, got %s", tt.expected, err)
			} else if err == nil && tt.expected == nil {
				// Check the value if there is no error
				for i, v := range value {
					if v != tt.doubledSlice[i] {
						t.Errorf("Expected value %d, got %d", tt.doubledSlice[i], v)
					}
				}
			}
		})
	}
}

func BenchmarkDoubledSlice(b *testing.B) {
	r := rand.New(rand.NewPCG(0, 12345)) // Fixed seed
	slice := make([]int, 100000)
	for i := range slice {
		slice[i] = r.IntN(100)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculator.DoubleSlice(slice)
	}
}

func BenchmarkDoubleSliceWithFunc(b *testing.B) {
	r := rand.New(rand.NewPCG(0, 12345))
	slice := make([]int, 100000)
	for i := range slice {
		slice[i] = r.IntN(100)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculator.DoubleSliceWithFunc(slice)
	}
}
