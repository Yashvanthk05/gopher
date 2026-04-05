package main

import (
	"fmt"
	"sync"
)

// go uses writer preference that is if writer and multiple readers are waiting preference is given to writer in order to avoid writer starvation

var data int

func read(mu *sync.RWMutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.RLock() // reader lock
	defer mu.RUnlock() // reader unlock
	fmt.Println(data)
}

func write(mu *sync.RWMutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock() // writer lock
	defer mu.Unlock() // writer unlock
	data++
}

func main() {
	var mu sync.RWMutex
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go read(&mu,&wg)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go write(&mu, &wg)
	}

	wg.Wait()

}