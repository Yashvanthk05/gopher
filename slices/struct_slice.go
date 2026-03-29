package main

import (
	"fmt"
)

type Student struct{
	name string
	regno int
}

func main(){
	// composite - create and initialize a data structure in one step, always requires ',' in newline of composite literal
	students := []Student{
		{"yashvanth",1589},
		{"kavin",1320},
	}

	fmt.Println(students)
}