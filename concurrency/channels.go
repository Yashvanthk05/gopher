package main

import "fmt"

// channel in Go is a typed communication pipe used by goroutines to send and receive values safely

// channel is like a pipe between goroutines where one goroutine puts data in and another goroutine takes data out and go runtime ensures synchronization (no race conditions)

// Channels synchronize goroutines by making them wait for each other by default

// Blocking by default = channel operations wait until the other side is ready

func main() {
	// unbuffered channels blocks until both sender and receiver are ready
	ch := make(chan int)

	go func(){
		// blocks until some receives
		ch <- 10
	}()

	// blocks until some sends
	fmt.Println(<-ch)

	// buffered channel - asynchoronous till buffer limit
	// sender is not blocked till buffer limit
	// receiver is not blocked till buffer is empty
	ch1 := make(chan int, 1)
	go func(){
		ch1 <- 1
		// blocked until 1 is received then 2 is pushed and received by main routine
		ch1 <- 2
	}()
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
}