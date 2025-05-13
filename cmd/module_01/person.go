package main

import (
	"fmt"

	"github.com/abdullinmm/grok-lean-go/internal/person"
)

func PersonMain() {
	p := person.Person{Name: "Marsel", Balance: 100.00}

	p.AddMoney(200.00)
	p.SpendMoney(50.00)

	fmt.Printf("Your balance %s = %.2f", p.Name, p.Balance)
}
