package slicecheck

import "errors"

var (
	// ErrSliceNil indicates that the slice is nil.
	ErrSliceNil = errors.New("Slice is nil")
	// ErrEmptySlice indicates that the slice is empty.
	ErrEmptySlice = errors.New("Slice is empty")
	// ErrInvalidType indicates that the slice contains an invalid type.
	ErrInvalidType = errors.New("Slice contains an invalid type")
)

// ValidateSlice validates a slice of any type.
func ValidateSlice[T any](slice []T) error {
	// check for nil slice
	if slice == nil {
		return ErrSliceNil
	}
	if len(slice) == 0 {
		return ErrEmptySlice
	}
	return nil
}
