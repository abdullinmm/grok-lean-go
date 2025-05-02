package calculator

import (
	"errors"
	"math"
)

var (
	ErrOverflow            = errors.New("Overflow when calculating the sum or difference")
	ErrOverflowAddition    = errors.New("Overflow during addition")
	ErrOverflowSubtraction = errors.New("Overflow during subtraction")
)

// Calculate returns the sum and difference of two numbers.
// If the result of the operation exceeds the maximum allowed float64 value,
// an error is returned.
func CalculateFloat(a, b float64) (sum, difference float64, err error) {

	sum = a + b
	difference = a - b

	if math.IsInf(sum, 0) || math.IsInf(difference, 0) {
		err = ErrOverflow
		return
	}
	return
}

// Calculate returns the sum and difference of two integers.
// If an overflow has occurred, an error is returned.
func CalculateInt(a, b int) (sum, difference int, err error) {

	sum = a + b
	difference = a - b

	// Checking overflow during addition
	if (a > 0 && b > 0 && sum < 0) || (a < 0 && b < 0 && sum > 0) {
		err = ErrOverflowAddition
		return
	}

	// Checking overflow during subtraction
	if (a >= 0 && b < 0 && difference < 0) || (a < 0 && b > 0 && difference > 0) {
		err = ErrOverflowSubtraction
		return
	}
	return
}
