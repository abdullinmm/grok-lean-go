package worker

import (
	"testing"

	"github.com/abdullinmm/grok-lean-go/internal/worker"
)

func BenchmarkAddNumbers(b *testing.B) {
	number := make([]int, 10000)

	for i := range len(number) {
		number[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		worker.AddNumbers(number)
	}
}

func BenchmarkAddNumbersMu(b *testing.B) {
	number := make([]int, 10000)

	for i := range len(number) {
		number[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		worker.AddNumberMu(number)
	}
}

func BenchmarkAddNumbersCh(b *testing.B) {
	number := make([]int, 10000)

	for i := range len(number) {
		number[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		worker.AddNumberCh(number)
	}
}

func BenchmarkAddNumbersCh1(b *testing.B) {
	number := make([]int, 10000)

	for i := range len(number) {
		number[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		worker.AddNumberCh1(number)
	}
}

func TestAddNumber(t *testing.T) {
	test := []struct {
		name     string
		slice    []int
		exptcted int
	}{
		{
			name:     "The sum of positive numbers",
			slice:    []int{1, 2, 3},
			exptcted: 6,
		},
		{
			name:     "The sum of negative numbers",
			slice:    []int{-1, 2, 3},
			exptcted: 4,
		},
		{
			name:     "Empty slice",
			slice:    []int{},
			exptcted: 0,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if v := worker.AddNumbers(tt.slice); v != tt.exptcted {
				t.Errorf("Expected %d, got %d", tt.exptcted, v)
			}
		})
	}
}
