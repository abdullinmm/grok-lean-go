package person

import (
	"errors"
	"testing"

	apperrors "github.com/abdullinmm/grok-lean-go/internal/errors"
)

func TestAddBalance(t *testing.T) {
	testCases := []struct {
		name     string
		person   *Person
		amount   float64
		balance  float64
		expected error
	}{
		{
			name:     "Valid deposit",
			person:   &Person{Name: "Marsel", Balance: 100.00},
			amount:   50,
			balance:  150,
			expected: nil,
		},
		{
			name:     "Negative amount",
			person:   &Person{Name: "Marsel", Balance: 100.00},
			amount:   -50.00,
			balance:  100.00,
			expected: &apperrors.NegativeAmountError{Amount: -50},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.person.AddMoney(tc.amount)
			if tc.expected != nil {
				// Check if both expected and actual errors are of type NegativeAmountError
				var expectedNeg, gotNeg *apperrors.NegativeAmountError
				if errors.As(tc.expected, &expectedNeg) && errors.As(err, &gotNeg) {
					// Compare Amount fields for deeper validation
					if gotNeg.Amount != expectedNeg.Amount {
						t.Errorf("expected amount %v, got %v", expectedNeg.Amount, gotNeg.Amount)
					}
				} else if err == nil {
					// Expected error but got nil
					t.Errorf("expected error %v, got nil", tc.expected)
				} else if err.Error() != tc.expected.Error() {
					// Unexpected error type or message
					t.Errorf("expected error %v, got %v", tc.expected, err)
				}
			} else {
				// No error expected
				if err != nil {
					t.Errorf("did not expect error, got: %v", err)
				}
				// Check if balance is as expected after operation
				if tc.balance != tc.person.Balance {
					t.Errorf("expected balance %.2f, got %.2f", tc.balance, tc.person.Balance)
				}
			}
		})
	}
}

func TestSpendBalance(t *testing.T) {
	testCases := []struct {
		name     string
		person   *Person
		amount   float64
		balance  float64
		expected error
	}{
		{
			name:     "Valid spend balance",
			person:   &Person{Name: "Marsel", Balance: 100.00},
			amount:   50,
			balance:  50,
			expected: nil,
		},
		{
			name:     "The amount exceeds the balance",
			person:   &Person{Name: "Marsel", Balance: 50.00},
			amount:   150.00,
			balance:  50.00,
			expected: &apperrors.InsufficientFundsError{Required: 150, Available: 50},
		},
		{
			name:     "Negative amount",
			person:   &Person{Name: "Marsel", Balance: 100.00},
			amount:   -50.00,
			balance:  100.00,
			expected: &apperrors.NegativeAmountError{Amount: -50},
		},
		{
			name:     "The amount is equal to the balance",
			person:   &Person{Name: "Marsel", Balance: 100.00},
			amount:   100.00,
			balance:  0.00,
			expected: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.person.SpendMoney(tc.amount)
			if tc.expected != nil {
				// Check both possible error types
				var expectedNeg, gotNeg *apperrors.NegativeAmountError
				var expectedFunds, gotFunds *apperrors.InsufficientFundsError
				if errors.As(tc.expected, &expectedNeg) && errors.As(err, &gotNeg) {
					// Compare Amount fields for NegativeAmountError
					if gotNeg.Amount != expectedNeg.Amount {
						t.Errorf("expected amount %v, got %v", expectedNeg.Amount, gotNeg.Amount)
					}
				} else if errors.As(tc.expected, &expectedFunds) && errors.As(err, &gotFunds) {
					// Compare Required and Available fields for InsufficientFundsError
					if gotFunds.Required != expectedFunds.Required || gotFunds.Available != expectedFunds.Available {
						t.Errorf("expected funds error Required %.2f Available %.2f, got Required %.2f Available %.2f", expectedFunds.Required, expectedFunds.Available, gotFunds.Required, gotFunds.Available)
					}
				} else if err == nil {
					// Expected error but got nil
					t.Errorf("expected error %v, got nil", tc.expected)
				} else if err.Error() != tc.expected.Error() {
					// Unexpected error type or message
					t.Errorf("expected error %v, got %v", tc.expected, err)
				}
			} else {
				// No error expected
				if err != nil {
					t.Errorf("did not expect error, got: %v", err)
				}
				// Check if balance is as expected after operation
				if tc.balance != tc.person.Balance {
					t.Errorf("expected balance %.2f, got %.2f", tc.balance, tc.person.Balance)
				}
			}
		})
	}
}
