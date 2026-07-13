package main

import "fmt"

func Display[T any] (val T) {
	fmt.Println("Variable Value:",val)
} 

type Number interface {
	int | float64
}

func Add[T Number] (a, b T) {
	fmt.Println(a+b)
}

func main() {
	Display[float64](3.144)
	Display[string]("Yashvanth")
	Display[struct{name string}](struct{name string}{name: "Yashvanth"})
	Add[float64](3.14,3.14)
	Add[int](3,2)
}