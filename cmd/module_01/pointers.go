package main

import (
	"fmt"

	"github.com/abdullinmm/grok-lean-go/internal/models"
)

func pointers() {
	person := &models.Person{}
	person.Name = "Masel"
	person.Age = 43

	fmt.Printf("User name befor update : %s. User age before update: %d\n", person.Name, person.Age)

	err := person.ValidatePerson()
	if err != nil {
		fmt.Printf("Invalid struct %v", err)
	}

	name := "Marat"
	age := 50

	person.UpdatePerson(name, age)
	fmt.Printf("User name after update: %s. User age after update: %d\n", person.Name, person.Age)
	s := []int{1, 2, 3}
	updateSlice(s)
	fmt.Println(s)
}

func updateSlice(s []int) {
	s[0] = 999
	s = append(s, 4)
}
