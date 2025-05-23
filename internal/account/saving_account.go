package account

import (
	"github.com/abdullinmm/grok-lean-go/internal/errors"
)

type SavingAccount struct {
	Balance float64
}

func (sa *SavingAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return &errors.NegativeAmountError{Amount: amount}
	}
	sa.Balance += amount
	return nil
}

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

func (sa *SavingAccount) GetBalance() float64 {
	return sa.Balance
}
