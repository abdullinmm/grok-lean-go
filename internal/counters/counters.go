package counters

import "sync"

// CounterChannel represents a channel-based counter.
type CounterChannel struct {
	readCh  chan uint64
	writeCh chan struct{}
}

// NewCounterChannel creates a new CounterChannel instance.
func NewCounterChannel() *CounterChannel {
	c := &CounterChannel{
		readCh:  make(chan uint64),
		writeCh: make(chan struct{}),
	}

	// Goroutine is the owner of the counter
	go func() {
		var count uint64 = 0
		for {
			select {
			case c.readCh <- count:
			case <-c.writeCh:
				count++
			}
		}
	}()
	return c
}

// Inc increments the counter.
func (c *CounterChannel) Inc() {
	c.writeCh <- struct{}{}
}

// Value returns the current value of the counter.
func (c *CounterChannel) Value() uint64 {
	return <-c.readCh
}

// CounterCh creates a counter channel and increments it by the specified number of workers.
func CounterCh(workers int) uint64 {
	c := NewCounterChannel()
	var wg sync.WaitGroup

	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()

	return c.Value()
}
