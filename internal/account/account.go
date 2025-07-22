package account

// Account represents a bank account.
type Account interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}
