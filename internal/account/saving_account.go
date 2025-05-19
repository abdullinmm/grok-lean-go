package account

import "github.com/abdullinmm/grok-lean-go/internal/person"

type SavingAccount struct {
	Balance float64
}

func (sa *SavingAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return person.ErrNegativeAmount
	}
	sa.Balance += amount
	return nil
}

func (sa *SavingAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return person.ErrNegativeAmount
	}
	if sa.Balance < amount {
		return person.ErrInsufficientFunds
	}
	sa.Balance -= amount
	return nil
}
