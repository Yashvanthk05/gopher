package main

import (
	"fmt"
	"sync"
)

func add(cnt *int, wg *sync.WaitGroup, mu *sync.RWMutex) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	*cnt++
}

func main() {
	var wg sync.WaitGroup
	var mu sync.RWMutex
	i := 0
	cnt := 0
	for i<10 {
		wg.Add(1)
		go add(&cnt, &wg, &mu)
		i++
	}
	wg.Wait()
	fmt.Println(cnt)
}