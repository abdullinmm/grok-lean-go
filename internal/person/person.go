package person

import (
	"github.com/abdullinmm/grok-lean-go/internal/errors"
)

type Person struct {
	Name    string
	Balance float64
}

// Adds money to the balance, returns an error if amount < 0
func (p *Person) AddMoney(amount float64) error {
	if amount < 0 {
		return &errors.NegativeAmountError{Amount: amount}
	}

	p.Balance += amount
	return nil
}

// Spends money, returns an error if amount < 0 or insufficient funds
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

func (p *Person) Withdraw(amount float64) error {
	return p.SpendMoney(amount)
}

func (p *Person) GetBalance() float64 {
	return p.Balance
}
