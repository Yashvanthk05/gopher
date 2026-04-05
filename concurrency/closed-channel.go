package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	// loop stops only when the channel is both closed and empty
	// val, ok := <-ch  // ok=true (value successfully received), ok=false (channel is closed and empty)
	// channel is closed and empty
	for v := range ch {
		fmt.Println(v)
	}
}