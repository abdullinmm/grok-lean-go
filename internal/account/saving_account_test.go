package account

import (
	"testing"

	"github.com/abdullinmm/grok-lean-go/internal/errors"
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
		{"Zero deposit", 50, 0, 50, &errors.NegativeAmountError{}},
		{"Negative deposit", 50, -10, 50, &errors.NegativeAmountError{}},
		{"Fractional deposit", 100.50, 0.50, 101.00, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sa := SavingAccount{Balance: tt.initial}
			err := sa.Deposit(tt.amount)

			if tt.wantErr != nil {
				assert.ErrorAs(t, err, &tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
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
		{"Insufficient funds", 50.50, 100, 50.50, &errors.InsufficientFundsError{}},
		{"Negative withdrawal", 100, -50, 100, &errors.InsufficientFundsError{}},
		{"Exact balance withdrawal", 75.75, 75.75, 0, nil},
		{"Fractional withdrawal", 100.99, 0.99, 100.00, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sa := SavingAccount{Balance: tt.initial}
			err := sa.Withdraw(tt.amount)

			if tt.wantErr != nil {
				assert.ErrorAs(t, err, &tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.wantBalance, sa.Balance)
		})
	}
}

// === Tests for Person (via Account interface) ===
func TestPerson_AccountInterface(t *testing.T) {
	tests := []struct {
		name        string
		setup       func() (Account, func() error)
		wantBalance float64
		wantErr     error
	}{
		{
			name: "Person Deposit",
			setup: func() (Account, func() error) {
				p := &person.Person{Balance: 0}
				return p, func() error { return p.Deposit(100) }
			},
			wantBalance: 100,
			wantErr:     nil,
		},
		{
			name: "Person Withdraw",
			setup: func() (Account, func() error) {
				p := &person.Person{Balance: 0}
				return p, func() error {
					if err := p.Deposit(100); err != nil {
						return err
					}
					return p.Withdraw(50)
				}
			},
			wantBalance: 50,
			wantErr:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc, op := tt.setup()
			err := op()

			if tt.wantErr != nil {
				assert.ErrorAs(t, err, &tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.wantBalance, acc.(*person.Person).Balance)
		})
	}
}

// === Combined tests ===
// func TestAccountOperations(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		account account.Account
// 	}{
// 		{"SavingAccount", SavingAccount{}},
// 		{"Person", &person.Person{}},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Initial deposit
// 			err := tt.account.Deposit(100)
// 			assert.NoError(t, err)

// 			// Withdrawal of part of funds
// 			err = tt.account.Withdraw(50)
// 			assert.NoError(t, err)

// 			// Attempt to withdraw more balance
// 			err = tt.account.Withdraw(60)
// 			expectedErr := &errors.InsufficientFundsError{
// 				Required:  60,
// 				Available: 50, // After 100 deposit and 50 withdrawal, the balance is 50
// 			}
// 			assert.ErrorAs(t, err, &expectedErr)
// 		})
// 	}
// }
