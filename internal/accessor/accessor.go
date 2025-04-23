package accessor

import (
	"errors"
)

var (
	ErrIndexOutOfBounds = errors.New("index out of range [%d] with length %d")
	ErrInvalidMapKey    = errors.New("invalid map key %s")
	ErrUnsupportedType  = errors.New("unsupported type: %T")
)

// Get safely retrieves an element from a slice by index.
func GetIndex(slice []int, index int) (int, error) {
	if index < 0 || index > len(slice) {
		return 0, ErrIndexOutOfBounds
	}
	return index, nil
}

// Get valid map key
func GetKey(m map[string]int, key string) (string, error) {
	if _, ok := m[key]; !ok {
		return "", ErrInvalidMapKey
	}
	return key, nil
}
