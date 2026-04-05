package main

import "fmt"

type A struct {
	name string
}

type B struct {
	A // composition over inheritance
}

func (a A) greet() {
	fmt.Println("Hello",a.name)
}

// embedded struct method override(shadowing)
func (b B) greet() {
	fmt.Println("Welcome",b.name)
}

func main() {
	a := A{name: "yashvanth"}
	a.greet()
	b := B{A{name: "yashvanth"}}
	b.greet()
	// accessing embedded method directly
	b.A.greet()
	// order of lookup for methods: first B methods then only embedded A methods
}