package main

import (
	"fmt"

	"github.com/abdullinmm/grok-lean-go/internal/account"
	"github.com/abdullinmm/grok-lean-go/internal/person"
)

func AccountMain() {
	// Initializing Person and SavingsAccount
	p := &person.Person{}
	sa := &account.SavingAccount{}

	// Slice of account
	accounts := []account.Account{p, sa}

	for _, acc := range accounts {
		fmt.Printf("Processing account: %T\n", acc)
		if err := acc.Deposit(100); err != nil {
			fmt.Println("Deposit Error:", err)
		}
		if err := acc.Withdraw(50); err != nil {
			fmt.Println("Withdraw Error:", err)
		}
	}

	// Print of balances
	fmt.Printf("Person Balance: %.2f\n", p.Balance)
	fmt.Printf("SavingsAccount Balance: %.2f\n", sa.Balance)
}
