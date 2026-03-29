package main

import (
	"fmt"
	"math"
)

type Shape interface{
	Area() float64
}

type Circle struct{
	radius float64
}

type Rect struct{
	width float64
	height float64
}

func (c Circle) Area() float64 {
	return math.Pi*c.radius*c.radius
}

func (r Rect) Area() float64 {
	return r.height*r.width
}

func main(){
	var rect Shape = Rect{10,20}
	//var cir Shape = Circle{10}

	// extracts the base type
	t1,ok := rect.(Circle)
	if ok {
		fmt.Println("It is a Circle",t1)
	}else{
		fmt.Println("It is a Rectangle")
	}

	// keep interfaces as small as possible to represent standard clean idea
	// interfaces should not be aware of base types eg. isFireTruck()
	// interfaces are not classes (even slimmer)
	// they don't have constructor and destructor; no inheritance its just signatures
	// type switches
	switch v := rect.(type) {
	case Rect:
		fmt.Println("It is a Rectangle",v)
	case Circle:
		fmt.Println("It is a Circle",v)
	default:
		fmt.Println("It is of unknown type")
	}
}