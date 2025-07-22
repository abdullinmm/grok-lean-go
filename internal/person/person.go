package person

import (
	"github.com/abdullinmm/grok-lean-go/internal/errors"
)

// Person represents a person with a name and balance.
type Person struct {
	Name    string
	Balance float64
}

// AddMoney adds the specified amount to the account balance.
func (p *Person) AddMoney(amount float64) error {
	if amount < 0 {
		return &errors.NegativeAmountError{Amount: amount}
	}

	p.Balance += amount
	return nil
}

// SpendMoney spends the specified amount from the account balance.
func (p *Person) SpendMoney(amount float64) error {
	if amount < 0 {
		return &errors.NegativeAmountError{Amount: amount}
	}
	if p.Balance < amount {
		return &errors.InsufficientFundsError{Required: amount, Available: p.Balance}
	}

	p.Balance -= amount
	return nil
}

// Deposit adds the specified amount to the account balance.
func (p *Person) Deposit(amount float64) (err error) {
	return p.AddMoney(amount)
}

// Withdraw spends the specified amount from the account balance.
func (p *Person) Withdraw(amount float64) error {
	return p.SpendMoney(amount)
}

// GetBalance returns the current balance of the account.
func (p *Person) GetBalance() float64 {
	return p.Balance
}
