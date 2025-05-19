package person

import "errors"

var (
	ErrNegativeAmount    = errors.New("amount must be positive")
	ErrInsufficientFunds = errors.New("Insufficient funds")
)

type Person struct {
	Name    string
	Balance float64
}

// adds money to the balance, returns an error if amount < 0
func (p *Person) AddMoney(amount float64) error {
	if amount < 0 {
		return ErrNegativeAmount
	}

	p.Balance += amount
	return nil
}

// spends money, returns an error if amount < 0 or insufficient funds
func (p *Person) SpendMoney(amount float64) error {
	if amount < 0 {
		return ErrNegativeAmount
	}
	if p.Balance < amount {
		return ErrInsufficientFunds
	}

	p.Balance -= amount
	return nil
}

func (p *Person) Deposit(amount float64) (err error) {
	return p.AddMoney(amount)
}

func (p *Person) Withdraw(amount float64) error {
	return p.SpendMoney(amount)
}
