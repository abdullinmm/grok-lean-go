package worker

import "fmt"

// ProccesChannel ....
func ProccesChannel(ch <-chan int) {
	for {
		value, ok := <-ch
		if !ok {
			// channel closed, leave the cycle
			fmt.Println("Channel close.")
			break
		}
		// Processing the resulting value
		fmt.Printf("Get value: %d\n", value)
	}
}
