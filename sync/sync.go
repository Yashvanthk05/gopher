package main

import (
	"fmt"
	"sync"
)

// sync waitgroup
// since main function does not wait for goroutines to finish
// goroutines maintains a counter and wait till the counter becomes zero

// sync mutex
// only one goroutine access a critical section at a time so this avoid the race condition that is unpredictable value


func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Worker",id,"Started")
	fmt.Println("Worker",id,"Finished")
}

func main() {
	var wg sync.WaitGroup
	n := 5
	wg.Add(n)
	for i := 0; i < n; i++ {
		go worker(i+1,&wg)
	}
	wg.Wait()
	fmt.Println("All workers finished")

	// without mutext this leads to final unpredictable value
	count := 0
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++
		}()
	}

	wg.Wait()
	// final wrong value due to race condition that two or more goroutines accessing the critical section at the same time
	fmt.Println("Without Mutex:",count)
	
	// with mutext
	var mu sync.Mutex
	count = 0
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock() // ensure the mutex is always unblocks incase of panic or early exits
			count++
		}()
	}

	wg.Wait()
	fmt.Println("With Mutex:",count)
}