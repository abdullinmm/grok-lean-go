package main

import (
	"errors"
	"fmt"

	"github.com/abdullinmm/grok-lean-go/internal/account"
	apperrors "github.com/abdullinmm/grok-lean-go/internal/errors"
	"github.com/abdullinmm/grok-lean-go/internal/person"
)

// AccountMain initializes Person and SavingsAccount, performs deposit and withdrawal operations, and prints balances.
func AccountMain() {
	// Initializing Person and SavingsAccount
	p := &person.Person{}
	sa := &account.SavingAccount{}

	var negErr *apperrors.NegativeAmountError

	// Slice of account
	accounts := []account.Account{p, sa}

	for _, acc := range accounts {
		fmt.Printf("Processing account: %T\n", acc)
		if err := acc.Deposit(100); err != nil {
			if errors.As(err, &negErr) {
				fmt.Println("Deposit Error:", err)
			}
		}
		if err := acc.Withdraw(50); err != nil {
			fmt.Println("Withdraw Error:", err)
		}
	}

	// Print of balances
	fmt.Printf("Person Balance: %.2f\n", p.Balance)
	fmt.Printf("SavingsAccount Balance: %.2f\n", sa.Balance)
}
