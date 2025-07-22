package main

import (
	"fmt"

	"github.com/abdullinmm/grok-lean-go/internal/worker"
)

func main() {

	sliceNumber := make([]int, 4)

	for i := range len(sliceNumber) {
		sliceNumber[i] = i
	}

	fmt.Printf("The sum of all numbers in the cut AddNumber: %d\n", worker.AddNumbers(sliceNumber))
	fmt.Printf("The sum of all numbers in the cut AddNumberMu: %d\n", worker.AddNumbers(sliceNumber))
	fmt.Printf("The sum of all numbers in the cut AddNumberCh: %d\n", worker.AddNumbers(sliceNumber))
}
