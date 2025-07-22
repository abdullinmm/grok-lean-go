package counters

import "sync"

// Structure for managing counter channels
type CounterChannel struct {
	readCh  chan uint64
	writeCh chan struct{}
}

// Constructor
func NewCounterChannel() *CounterChannel {
	c := &CounterChannel{
		readCh:  make(chan uint64),
		writeCh: make(chan struct{}),
	}

	// Gorutina is the owner of the counter
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

func (c *CounterChannel) Inc() {
	c.writeCh <- struct{}{}
}

func (c *CounterChannel) Value() uint64 {
	return <-c.readCh
}

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
