package main

import (
	"fmt"

	"github.com/abdullinmm/grok-lean-go/internal/calculator"
)

func main() {
	name := "Alice"
	s := []int{2, 4, 7}
	num := 4

	average, err := calculator.CalculateAverage(s)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	printResult(name, s, average, num, calculator.IsEven(num))

	// slice := []int{1}

	// fmt.Println("Slice element:", slice[1])

	m := make(map[string]int)
	m["Marsel"] = 42
	value, ok := m["Marat"]
	if !ok {
		fmt.Println("Not valid key for map")
		return
	}
	fmt.Println("map value:", value)
}

func printResult(name string, grades []int, average float64, num int, isEven bool) {
	fmt.Printf("Hello, %s! Your grades: %v. Average grade: %.2f\n", name, grades, average)
	fmt.Printf("Number %d is even: %t\n", num, isEven)
}
