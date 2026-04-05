package main

// go promotes composition over classical inheritance
// Building complex types by combining smaller, reusable types

type Engine struct {
}

type Car struct {
	Engine
}

// Go’s Philosophy: “has-a” instead of “is-a”
// inheritance: Car is an Engine
// composition: Car has an Engine

/*
Reason why go prefers composition
- loose coupling, inheritance has a tight coupling child depends heavily on parent
- composition + interface = power
*/
func main() {

}