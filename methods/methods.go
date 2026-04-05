package main

import "fmt"

// named type Int
type Int int

// difference between a function and method is a receiver
// method - function attached to a type using a receiver

type Person struct {
	name string
	age int
}

// p Person is the receiver
func (p Person) greet() string {
	return "Hello " + p.name + "!"
}

// value receiver - works on a copy if p.name = "deepak" original will not be affected
func (p Person) modifyr() {
	p.name = "deepak"
} 

func (p *Person) modifyp() {
	p.name = "deepak"
}

// method on non-struct type
func (num Int) double() int {
	return int(num*2)
}

func main(){
	p := Person{name: "yashvanth",age: 21}
	fmt.Println(p.greet())

	// value receiver
	a := p
	a.modifyr()
	fmt.Println(a)

	// pointer receiver
	b := p
	b.modifyp()
	fmt.Println(b)

	// When to use pointer receiver — mutation, large structs (Avoids copying huge data every time)

	// Methods on non-struct types
	// named types based int, string etc..
	c := Int(10)
	fmt.Println(c.double())

	// 1. Method value - creates a function with receiver already fixed
	f1 := p.greet 
	fmt.Println(f1())
	// this is equivalent to
	// f1 := func(){
	// 	p.greet()
	// }

	// 2. Method Expression - create a regular function with receiver being passed explicitly
	f2 := Person.greet
	fmt.Println(f2(p))

	// method expression with pointer receiver
	f3 := (*Person).modifyp
	f3(&p)
	fmt.Println(p)
}