package main

import "fmt"

func conversion(converter func(int) int, x int) int {
	return converter(x)
}

func double(x int) int {
	return x*x
}

func main(){
	x := conversion(double, 2)
	fmt.Println(x)
	// anonymous function - function without a name
	a := 2
	b := 3

	// asssign to a variable
	c := func(a, b int) int {
		return a+b
	}

	fmt.Println(c(a,b))

	// immediately invoked  function expression
	// anonymous function that is executed immediately after it is defined
	// useful for scoping variables temporarily
	// avoid creating a separate named function
	// one time execution
	res := func(a int, b int) int {
		return a+b
	}(a,b)
	fmt.Println(res)

	// closure - capture outer variables
	func(){
		fmt.Println(a)
	}()
}