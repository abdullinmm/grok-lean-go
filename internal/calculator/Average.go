package calculator

import "errors"

var (
	errEmptySlice = errors.New("Slice is Empty")
)

// CalculateAverage calculates the average of a slice of integers.
func CalculateAverage(s []int) (float64, error) {
	if len(s) == 0 {
		return 0, errEmptySlice
	}
	sum := 0
	for _, v := range s {
		sum += v
	}
	return float64(sum) / float64(len(s)), nil
}
