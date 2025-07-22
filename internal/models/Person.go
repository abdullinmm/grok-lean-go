package models

import (
	"errors"
	"strings"
)

var (
	// ErrValidateName is returned when the name is invalid.
	ErrValidateName = errors.New("Invalid user name.")
	// ErrValidateAge is returned when the age is invalid.
	ErrValidateAge = errors.New("Age cannot be negative")
)

// Person represents a person with a name and age.
type Person struct {
	Name string
	Age  int
}

// ValidatePerson validates the Person struct.
func (p *Person) ValidatePerson() error {

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

// UpdatePerson updates the Person struct.
func (p *Person) UpdatePerson(name string, age int) {

	if p.Name != name {
		p.Name = name
	}

	if p.Age != age {
		p.Age = age
	}
}
