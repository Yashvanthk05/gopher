package main

import (
	"fmt"
	"math"
)

func Add(x, y int) (result int) {
	// named return value 'result' is initialized to zero by default
	// naked return statement returns the current value of 'result'
	// this is not recommended for larger functions as it can make the code less clear
	result = x + y
	return
}

func Double(x int) (result int) {
	// implicitly returning the named return value 'result' after modifying it
	result = int(math.Pow(float64(x),2))
	return
}

func Name(a, b string) (string, string) {
	// explicitly returning multiple values
	return a, b
}

func GuardClause(){
	// guard clause means that we check for a condition at the beginning of a function and return early if the condition is not met
	// this helps to reduce nesting and improve readability
	// for example, if we want to check if a number is negative before performing some operation on it, we can use a guard clause
	num := -5
	if num < 0 {
		fmt.Println("Negative number, returning early.")
		return
	}
	fmt.Println("Positive number, performing operation.")
}

func main(){
	// ignoring retrn value using blank identifier
	a,_ := Name("Yashvanth", "Karunakaran")
	fmt.Println(a)
	fmt.Println(Add(3, 4))
	fmt.Println(Double(5))
}