package slicecheck

import "errors"

var (
	ErrSliceNil   = errors.New("Slice is nil")
	ErrEmptySliсe = errors.New("Slice is empty")
)

// Generalized function
func ValidateSlice[T any](slice []T) error {
	// check for nil slice
	if slice == nil {
		return ErrSliceNil
	}
	if len(slice) == 0 {
		return ErrEmptySliсe
	}
	return nil
}
