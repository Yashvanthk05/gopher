package main

import "fmt"

func main(){
	var a int = 10
	// * - deference operator, & - address of operator
	var b *int = &a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)
	*b = 1
	fmt.Println(a,b,*b)
	
	// new(T) — allocate zeroed value, return pointer
	c := new(int)
	fmt.Println(c,*c)

	// No pointer arithmetic in Go like adding, subtracting or incrementing pointers is not allowed
	// c += 1 - not allowed - compiler time error
	// arbitary pointer movement would break GC tracking
	// use indexing in slices/arr instead of pointer movement
	// pointer movement possible with unsafe
	// Go disallows pointer arithmetic to ensure memory safety, simplicity, and compatibility with garbage collection.
	// When to use pointers — mutation, large structs, optional values

	// Nil pointer — dereferencing causes panic
	var p *int
	fmt.Println(p)
	// fmt.Println(*p) - causes panic
	if p!=nil {
		fmt.Println(*p)
	}
}