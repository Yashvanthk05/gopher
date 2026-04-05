package main

import (
	"fmt"
	"sync"
)

// sync.Once in Go is a synchronization primitive that guarantees a function is executed exactly one time, even in the presence of multiple concurrent goroutines.

/*
Multiple goroutines may call once.Do(f)
Only one goroutine executes f
All others wait until f completes
After first execution then subsequent calls do nothing
*/

// used for lazy initialization, loading config, initializing db connections

func initialize() {
	fmt.Println("Initialized")
}

func worker(wg *sync.WaitGroup, once *sync.Once) {
	defer wg.Done()
	once.Do(initialize)
	fmt.Println("Worker called")
}

func main() {
	var wg sync.WaitGroup
	var once sync.Once

	for i:=0;i<5;i++ {
		wg.Add(1)
		go worker(&wg,&once)
	}

	wg.Wait()
}