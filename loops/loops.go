package main

import (
	"fmt"
)

//omitting condition
func threshold(a int, thres int) int {
	sum := 0
	for i := 0; ; i++ {
		sum += a
		if sum == thres { return i }
	}
}

func main(){
	// each section of the for signature is optional
	// normal
	for i := 0; i < 5; i++ {
		fmt.Print(i," ")
	}
	fmt.Println()
	// omitting condition
	fmt.Println("Number of times:",threshold(2,10))
	// no while loop in go, for loop behave as while
	i := 0
	sum := 0
	for i <= 10 {
		sum += i
		i += 1
	}
	fmt.Println("Sum of First 10 natural numbers:",sum)
	// logical and - &&
	// logical or - ||
	// guard clause pattern within loops - continue and break
	// continue - skips the current iteration
	// break - stop the entire iteration
	// range just as in python

	fmt.Println("Range:")
	nums := []int{1,2,3,4,5,6}
	for idx, val := range nums {
		fmt.Printf("nums[%d]=%d; ",idx,val)
	}
	fmt.Println()
}