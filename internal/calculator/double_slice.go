package calculator

import (
	"github.com/abdullinmm/grok-lean-go/internal/slicecheck"
)

func DoubleSlice(slice []int) ([]int, error) {

	// check the slice for nil and for emptiness
	if err := slicecheck.ValidateSlice(&slice); err != nil {
		return nil, err
	}

	double := make([]int, len(slice))
	for i, v := range slice {
		double[i] = v * 2
	}

	return double, nil
}

func DoubleSliceWithFunc(slice []int) ([]int, error) {

	// check the slice for nil and for emptiness
	if err := slicecheck.ValidateSlice(&slice); err != nil {
		return nil, err
	}

	// the anonymous function doubles the number
	apply := func(n int) int { return n * 2 }

	doubleSlice := make([]int, len(slice))
	for i, v := range slice {
		doubleSlice[i] = apply(v)
	}

	return doubleSlice, nil
}
