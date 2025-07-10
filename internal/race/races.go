package race

import (
	"sync"
)

func Counter(workers int) (count int) {

	if workers < 1 {
		return count
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}

	wg.Wait()

	return count
}
