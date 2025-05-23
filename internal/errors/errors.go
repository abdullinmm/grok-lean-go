package errors

import "fmt"

type InsufficientFundsError struct {
	Required  float64
	Available float64
}

func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("insufficient funds: required %.2f, available %.2f", e.Required, e.Available)
}

type NegativeAmountError struct {
	Amount float64
}

func (e *NegativeAmountError) Error() string {
	return fmt.Sprintf("amount can't be negative amount %.2f", e.Amount)
}
