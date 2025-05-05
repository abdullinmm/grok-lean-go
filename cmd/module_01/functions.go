package main

import (
	"fmt"

	"github.com/abdullinmm/grok-lean-go/internal/calculator"
)

func Functions() {

	a := 2
	b := 3

	sum, diff, err := calculator.Calculate(a, b)
	if err != nil {
		return
	}

	fmt.Printf("sum = %d, diff = %d\n", sum, diff)

	// an anonymous function that doubles the result of the sum.
	double := func() int {
		return sum * 2
	}

	fmt.Printf("double = %d", double())
}
