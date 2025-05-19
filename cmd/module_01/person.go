package main

import (
	"fmt"

	"github.com/abdullinmm/grok-lean-go/internal/person"
	"github.com/abdullinmm/grok-lean-go/internal/slicecheck"
)

func PersonMain() {
	p := person.Person{Name: "Marsel", Balance: 100.00}
	p1 := person.Person{Name: "Emil", Balance: 10.00}
	p2 := person.Person{Name: "Marat", Balance: 110.00}

	slice := []person.Person{p, p1, p2}

	nameSlice, err := CheckBalance(slice)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("the names of all people whose balance is more than 100.00: %v\n", nameSlice)

	err = p.Deposit(200.00)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = p.Withdraw(50.00)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Your balance %s = %.2f\n", p.Name, p.Balance)
}

func CheckBalance(p []person.Person) (slice []string, err error) {

	// check nil and empty slice
	if err = slicecheck.ValidateSlice(p); err != nil {
		return
	}

	for _, v := range p {
		if v.Balance <= 100 {
			slice = append(slice, v.Name)
		}
	}
	return
}
