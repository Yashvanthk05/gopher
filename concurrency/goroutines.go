package main

import (
	"fmt"
	"time"
	"runtime"
)

// go routine is a lightweight thread becoz it each goroutine starts with a small stack size of ~1kB and grows dynamically whereas traditional OS thread has fixed stack size of ~1MB so it is very expensive


// go prefers sharing memory by communicating over communicating by shared memory
// memory is shared by channels

func display() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond*100)
		fmt.Println("inside display:",i)
	}
}

func main(){
	go display()

	// this loop will finish and if main routines finishes all other routines will be stopped
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(i)
	}

	fmt.Println("No of CPU cores go is allowed to use:",runtime.GOMAXPROCS(0))
}