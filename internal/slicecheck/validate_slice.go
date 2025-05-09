package slicecheck

import "errors"

var (
	ErrSliceNil   = errors.New("Slice is nil")
	ErrEmptySlise = errors.New("Slice is empty")
)

func ValidateSlice(slice *[]int) error {
	// check for nil slice
	if *slice == nil {
		return ErrSliceNil
	}
	// check slice len
	if len(*slice) <= 0 {
		return ErrEmptySlise
	}
	return nil
}
