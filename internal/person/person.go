package person

import "errors"

var (
	ErrNegativeAmount    = errors.New("Negative amount")
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
	if amount < 0 || p.Balance < amount {
		return ErrInsufficientFunds
	}

	p.Balance -= amount
	return nil
}
