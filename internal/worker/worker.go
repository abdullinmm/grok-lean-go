package worker

import (
	"fmt"
	"sync"

	"github.com/abdullinmm/grok-lean-go/internal/slicecheck"
)

func ProcessNumber(number []int, workers int) []int {
	if err := slicecheck.ValidateSlice(number); err != nil {
		fmt.Println("Slice error:", err)
		return []int{}
	}

	if workers < len(number) || workers < 1 {
		workers = len(number)
	}

	results := make([]int, len(number))
	ch := make(chan struct {
		index int
		value int
	}, len(number))
	jobs := make(chan int, len(number))

	go func() {
		for i := range number {
			jobs <- i
		}
		close(jobs)
	}()

	var wg sync.WaitGroup
	wg.Add(workers)

	for range workers {
		go func() {
			defer wg.Done()
			for idx := range jobs {
				ch <- struct {
					index int
					value int
				}{idx, number[idx] * 2}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		results[res.index] = res.value
	}

	return results
}
