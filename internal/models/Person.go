package models

import (
	"errors"
	"strings"
)

var (
	ErrValidateName = errors.New("Invalit user name.")
	ErrValidateAge  = errors.New("Age cannot be negative")
)

type Person struct {
	Name string
	Age  int
}

// User validation method (uses a pointer to avoid copying)
func (p *Person) Validate() error {

	// Name validation
	if strings.TrimSpace(p.Name) == "" || len(p.Name) < 2 {
		return ErrValidateName
	}

	// Age validation
	if p.Age <= 0 {
		return ErrValidateAge
	}
	return nil
}

// Updates the Name and Age fields(uses a pointer to avoid copying)
func (p *Person) UpdatePerson(name string, age int) {

	if p.Name != name {
		p.Name = name
	}

	if p.Age != age {
		p.Age = age
	}
}
