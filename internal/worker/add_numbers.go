package worker

import (
	"runtime"
	"sync"
)

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

func AddNumberMu(number []int) int {
	if number == nil {
		return 0
	}

	result := 0

	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(len(number))

	for v := range len(number) {
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

func AddNumberCh(number []int) int {
	if len(number) == 0 {
		return 0
	}
	//workers := len(number)
	workers := runtime.NumCPU()
	results := make(chan int, workers)
	chunkSize := (len(number) + workers - 1) / workers

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(number) {
			end = len(number)
		}
		wg.Add(1)
		go func(nums []int) {
			defer wg.Done()
			sum := 0
			for _, v := range nums {
				sum += v
			}
			results <- sum
		}(number[start:end])
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

	// Каскадная редукция
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
