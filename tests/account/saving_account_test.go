package account

import (
	"testing"

	"github.com/abdullinmm/grok-lean-go/internal/account"
	"github.com/abdullinmm/grok-lean-go/internal/person"
	"github.com/stretchr/testify/assert"
)

// === Tests for SavingAccount ===
func TestSavingAccount_Deposit(t *testing.T) {
	tests := []struct {
		name        string
		initial     float64
		amount      float64
		wantBalance float64
		wantErr     error
	}{
		{"Positive deposit", 0, 100.50, 100.50, nil},
		{"Zero deposit", 50, 0, 50, person.ErrNegativeAmount},
		{"Negative deposit", 50, -10, 50, person.ErrNegativeAmount},
		{"Fractional deposit", 100.50, 0.50, 101.00, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sa := &account.SavingAccount{Balance: tt.initial}
			err := sa.Deposit(tt.amount)

			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.wantBalance, sa.Balance)
		})
	}
}

func TestSavingAccount_Withdraw(t *testing.T) {
	tests := []struct {
		name        string
		initial     float64
		amount      float64
		wantBalance float64
		wantErr     error
	}{
		{"Successful withdrawal", 100.50, 50.25, 50.25, nil},
		{"Insufficient funds", 50.50, 100, 50.50, person.ErrInsufficientFunds},
		{"Negative withdrawal", 100, -50, 100, person.ErrNegativeAmount},
		{"Exact balance withdrawal", 75.75, 75.75, 0, nil},
		{"Fractional withdrawal", 100.99, 0.99, 100.00, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sa := &account.SavingAccount{Balance: tt.initial}
			err := sa.Withdraw(tt.amount)

			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.wantBalance, sa.Balance)
		})
	}
}

// === Tests for Person (via Account interface) ===
func TestPerson_AccountInterface(t *testing.T) {
	p := &person.Person{Balance: 0}

	tests := []struct {
		name        string
		account     account.Account
		operation   func() error
		wantBalance float64
		wantErr     error
	}{
		{
			name:    "Person Deposit",
			account: p,
			operation: func() error {
				return p.Deposit(100)
			},
			wantBalance: 100,
			wantErr:     nil,
		},
		{
			name:    "Person Withdraw",
			account: p,
			operation: func() error {
				p.Balance = 100
				return p.Withdraw(50)
			},
			wantBalance: 50,
			wantErr:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.operation()

			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.wantBalance, tt.account.(*person.Person).Balance)
		})
	}
}

// === Combined tests ===
func TestAccountOperations(t *testing.T) {
	tests := []struct {
		name    string
		account account.Account
	}{
		{"SavingAccount", &account.SavingAccount{}},
		{"Person", &person.Person{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initial deposit
			err := tt.account.Deposit(100)
			assert.NoError(t, err)

			// Withdrawal of part of funds
			err = tt.account.Withdraw(50)
			assert.NoError(t, err)

			// Attempt to withdraw more balance
			err = tt.account.Withdraw(60)
			assert.ErrorIs(t, err, person.ErrInsufficientFunds)
		})
	}
}
