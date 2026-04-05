package main

import "fmt"

func main(){
	const pi = 3.14 // untyped constant - go will infer the type based on how we will use it
	const name string = "Go" // typed constant

	// iota is counter, auto increments values in const blocks
	const (
		a = iota
		b
		c
	)

	fmt.Println(a,b,c)

	// const cannot store runtime values like time.Now() and complex structure like slices

	const (
		Read = 1 << iota   // 1 (001)
		Write              // 2 (010)
		Execute            // 4 (100)
	)

	fmt.Println(Read,Write,Execute)

	const (
		A = iota  // 0
		_         // skip 1
		B         // 2
	)

	fmt.Println(A,B)
}