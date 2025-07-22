package person

import (
	"errors"
	"testing"

	apperrors "github.com/abdullinmm/grok-lean-go/internal/errors"
)

func TestAddBalance(t *testing.T) {
	test := []struct {
		name     string
		person   *Person
		amount   float64
		balance  float64
		expected error
	}{
		{
			name:     "Valide spend balance",
			person:   &Person{Name: "Marsel", Balance: 100.00},
			amount:   50,
			balance:  150,
			expected: nil,
		},
		{
			name:     "Negative amount",
			person:   &Person{Name: "Marsel", Balance: 100.00},
			amount:   -50.00,
			balance:  00.00,
			expected: &apperrors.NegativeAmountError{Amount: -10},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.person.AddMoney(tt.amount)
			if !errors.Is(err, tt.expected) {
				t.Errorf("expected %s got %s", tt.expected, err)
			} else if err == nil && tt.expected == nil {
				if tt.balance != tt.person.Balance {
					t.Errorf("expected %.2f got %.2f", tt.balance, tt.person.Balance)
				}
			}
		})
	}
}

func TestSpendBalance(t *testing.T) {
	test := []struct {
		name     string
		person   *Person
		amount   float64
		balance  float64
		expected error
	}{
		{
			name:     "Valide spend balance",
			person:   &Person{Name: "Marsel", Balance: 100.00},
			amount:   50,
			balance:  50,
			expected: nil,
		},
		{
			name:     "The amount exceeds the balance",
			person:   &Person{Name: "Marsel", Balance: 50.00},
			amount:   150.00,
			balance:  00.00,
			expected: &apperrors.InsufficientFundsError{},
		},
		{
			name:     "Negative amount",
			person:   &Person{Name: "Marsel", Balance: 100.00},
			amount:   -50.00,
			balance:  00.00,
			expected: &apperrors.InsufficientFundsError{},
		},
		{
			name:     "The amount is equal to the balance",
			person:   &Person{Name: "Marsel", Balance: 100.00},
			amount:   100.00,
			balance:  00.00,
			expected: nil,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.person.SpendMoney(tt.amount)
			if !errors.Is(err, tt.expected) {
				t.Errorf("Expected %s got %s", err, tt.expected)
			} else if err == nil && tt.expected == nil {
				if tt.balance != tt.person.Balance {
					t.Errorf("Expected %.2f got %.2f", tt.balance, tt.person.Balance)
				}
			}
		})
	}
}
