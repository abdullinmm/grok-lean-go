package worker

import (
	"errors"
	"testing"

	"github.com/abdullinmm/grok-lean-go/internal/slicecheck"
	"github.com/abdullinmm/grok-lean-go/internal/worker"
)

func TestProccesNumber(t *testing.T) {
	test := []struct {
		name        string
		slice       []int
		doubleSlice []int
		exptcted    error
	}{
		{
			name:        "Valid test",
			slice:       []int{1, 2, 3, 4},
			doubleSlice: []int{2, 4, 6, 8},
			exptcted:    nil,
		},
		{
			name:     "Invalid test",
			exptcted: slicecheck.ErrSliceNil,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			value, err := worker.ProcessNumber(tt.slice, 4)
			if !errors.Is(err, tt.exptcted) {
				t.Errorf("Expected %s, got %s", tt.exptcted, err)
			} else if err == nil && tt.exptcted == nil {
				// Check the value if there is no error
				for i, v := range value {
					if v != tt.doubleSlice[i] {
						t.Errorf("Expected value %d, got %d", tt.doubleSlice[i], v)
					}
				}
			}
		})
	}
}
