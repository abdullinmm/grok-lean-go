package main

import (
	"fmt"

	"github.com/abdullinmm/grok-lean-go/internal/accessor"
	"github.com/abdullinmm/grok-lean-go/internal/calculator"
	"github.com/abdullinmm/grok-lean-go/internal/counters"
	"github.com/abdullinmm/grok-lean-go/internal/mathutils"
	"github.com/abdullinmm/grok-lean-go/internal/palindrome"
	"github.com/abdullinmm/grok-lean-go/internal/race"
	"github.com/abdullinmm/grok-lean-go/internal/worker"
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

	slice := []int{1}
	index := 0

	if _, err := accessor.GetIndex(slice, index); err != nil {
		fmt.Printf(fmt.Sprint(err), index, len(slice))
		return
	}

	m := make(map[string]int)
	key := "Marsel"
	m[key] = 42

	if _, err := accessor.GetKey(m, key); err != nil {
		fmt.Printf(fmt.Sprint(err), key)
		return
	}

	pointers()
	Functions()
	Slices()

	a := 10
	b := 5

	fmt.Printf("Find max value : %d\n", mathutils.Max(a, b))

	PersonMain()
	AccountMain()

	s1 := []int{2, 4, 7}

	if v, err := worker.ProcessNumber(s1, 5); err != nil {
		fmt.Printf("%v", err)
	} else {
		fmt.Printf("%v\n", v)
	}

	fmt.Printf("%v\n", race.Counter(60))
	// fmt.Printf("%v\n", race.CounterCh(60))
	fmt.Printf("%v\n", counters.CounterCh(1000))

	str := "NAN"
	if err := palindrome.Palindrome(str); err != nil {
		fmt.Printf("String %s is not palindrome", str)
	} else {
		fmt.Printf("String %s is palindrome", str)
	}
}

func printResult(name string, grades []int, average float64, num int, isEven bool) {
	fmt.Printf("Hello, %s! Your grades: %v. Average grade: %.2f\n", name, grades, average)
	fmt.Printf("Number %d is even: %t\n", num, isEven)
}

// Describe ....
func Describe(a any) string {
	switch val := a.(type) {
	default:
		return fmt.Sprintf("%T: %v", val, val)
	}
}
