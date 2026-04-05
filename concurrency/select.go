package main

import (
	"fmt"
	"time"
)

// select lets a goroutine wait on multiple channel operations and execute whichever is ready first

// Wait on many channels, and pick the one that is ready

/*
First ready case executes
- Go checks all cases
- If multiple are ready then one is chosen randomly
- If none are ready then it blocks (waits)
*/

// default case - non blocking (check and move on)

func worker (done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("Stopping worker")
			return
		default:
			fmt.Println("Working")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func closedChannel() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	for {
		select {
		case v, ok := <-ch1:
			// ch1 closed and empty
			if !ok {
				ch1 = nil // select will ignore nil channel
				continue
			}
			fmt.Println(v)
		case v, ok := <-ch2:
			// ch2 closed and empty
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Println(v)
		if ch1 == nil && ch2 == nil {
			break
		}
		}
	}
}
func main() {

	ch := make(chan int)

	// timeout pattern
	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(1*time.Second):
		fmt.Println("timeout")
	}

	// done/cancellation pattern
	// done channel used to tell goroutines to stop, stop signal sent to multiple running goroutines

	done := make(chan struct{})

	go worker(done)

	time.Sleep(2 * time.Second)

	close(done)

	time.Sleep(1 * time.Second)
}