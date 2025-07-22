package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/abdullinmm/grok-lean-go/internal/calculator"
)

// Slices
func Slices() {

	// Initializing the local generator
	r := rand.New((rand.NewPCG(0, uint64(time.Now().UnixNano()))))

	// create empty slice
	slice := make([]int, 5)

	// populate slice random values
	for i := range slice {
		slice[i] = r.IntN(100)
	}

	doubledSlice, err := calculator.DoubleSlice(slice)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	fmt.Printf("Slice values befor doubled: %v\nSlice values after doubled: %v\n", slice, doubledSlice)

	doubledSliceWithFunc, err := calculator.DoubleSlice(slice)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	fmt.Printf("Slice values befor doubled: %v\nSlice values after doubled: %v\n", slice, doubledSliceWithFunc)

	var sl []int

	calculator.DoubleSlice(sl)
}
