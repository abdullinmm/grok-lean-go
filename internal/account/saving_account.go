package account

import (
	"github.com/abdullinmm/grok-lean-go/internal/errors"
)

// SavingAccount represents a saving account with a balance.
type SavingAccount struct {
	Balance float64
}

// Deposit adds the specified amount to the account balance.
func (sa *SavingAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return &errors.NegativeAmountError{Amount: amount}
	}
	sa.Balance += amount
	return nil
}

// Withdraw subtracts the specified amount from the account balance.
func (sa *SavingAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return &errors.NegativeAmountError{Amount: amount}
	}
	if sa.Balance < amount {
		return &errors.InsufficientFundsError{Required: amount, Available: sa.Balance}
	}
	sa.Balance -= amount
	return nil
}

// GetBalance returns the current balance of the account.
func (sa *SavingAccount) GetBalance() float64 {
	return sa.Balance
}
