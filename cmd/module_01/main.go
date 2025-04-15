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

	fmt.Printf("Привет, %s! Твои оценки: %v. Средняя оценка: %.2f.", name, s, average)
	fmt.Printf("\nЧисло %d является четным: %t", num, calculator.IsEven(num))
}
