package errors

import "fmt"

// InsufficientFundsError represents an error when there are insufficient funds.
type InsufficientFundsError struct {
	Required  float64
	Available float64
}

// Error returns a string representation of the error.
func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("insufficient funds: required %.2f, available %.2f", e.Required, e.Available)
}

// NegativeAmountError represents an error when an amount is negative.
type NegativeAmountError struct {
	Amount float64
}

// Error returns a string representation of the error.
func (e *NegativeAmountError) Error() string {
	return fmt.Sprintf("amount can't be negative amount %.2f", e.Amount)
}
