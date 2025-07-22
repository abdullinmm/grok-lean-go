package calculator

import (
	"errors"
	"math"
)

var (
	// ErrOverflow indicates that an overflow has occurred.
	ErrOverflow = errors.New("Overflow when calculating the sum or difference")
	// ErrOverflowAddition indicates that an overflow has occurred during addition.
	ErrOverflowAddition = errors.New("Overflow during addition")
	// ErrOverflowSubtraction indicates that an overflow has occurred during subtraction.
	ErrOverflowSubtraction = errors.New("Overflow during subtraction")
	// ErrNumbersNegative indicates that the numbers cannot be negative.
	ErrNumbersNegative = errors.New("Numbers cannot be negative")
)

// CalculateFloat takes two float64 numbers and returns their sum and difference.
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

// CalculateInt takes two int numbers and returns their sum and difference.
// If an overflow has occurred during addition or subtraction, an error is returned.
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

// Calculate takes two int numbers and returns their sum and difference.
// If the numbers are negative, an error is returned.
func Calculate(a, b int) (sum, diff int, err error) {

	if (a < 0) || (b < 0) {
		err = ErrNumbersNegative
		return
	}

	sum = a + b
	diff = a - b

	return
}
