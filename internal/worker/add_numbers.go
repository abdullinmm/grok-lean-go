package worker

import (
	"runtime"
	"sync"
)

// AddNumbers sums all elements in the provided slice.
func AddNumbers(number []int) int {
	if number == nil {
		return 0
	}

	results := 0

	for _, v := range number {
		results = results + v
	}

	return results
}

// AddNumberMu sums all elements using a mutex for thread safety.
func AddNumberMu(numbers []int) int {
	if numbers == nil {
		return 0
	}

	result := 0

	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(len(numbers))

	for _, v := range numbers {
		go func() {
			defer wg.Done()
			mu.Lock()
			result = result + v
			mu.Unlock()
		}()
	}

	wg.Wait()

	return result
}

// AddNumberCh sums all elements in parallel using channels.
// The slice is split into chunks and each chunk is summed concurrently.
func AddNumberCh(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	workers := runtime.NumCPU()
	results := make(chan int, workers)
	chunkSize := (len(numbers) + workers - 1) / workers

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(numbers) {
			end = len(numbers)
		}
		wg.Add(1)
		go func(nums []int) {
			defer wg.Done()
			sum := 0
			for _, v := range nums {
				sum += v
			}
			results <- sum
		}(numbers[start:end])
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	total := 0
	for sum := range results {
		total += sum
	}
	return total
}

// AddNumberCh1 sums all elements concurrently and reduces partial results using a cascading reduction.
func AddNumberCh1(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	workers := runtime.NumCPU()
	chunkSize := (len(numbers) + workers - 1) / workers
	results := make(chan int, workers)
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(numbers) {
			end = len(numbers)
		}
		if start >= end {
			continue
		}
		wg.Add(1)
		go func(nums []int) {
			defer wg.Done()
			sum := 0
			for _, v := range nums {
				sum += v
			}
			results <- sum
		}(numbers[start:end])
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// cascading reduction
	partialSums := []int{}
	for sum := range results {
		partialSums = append(partialSums, sum)
	}
	for len(partialSums) > 1 {
		next := []int{}
		for i := 0; i < len(partialSums); i += 2 {
			if i+1 < len(partialSums) {
				next = append(next, partialSums[i]+partialSums[i+1])
			} else {
				next = append(next, partialSums[i])
			}
		}
		partialSums = next
	}
	return partialSums[0]
}
