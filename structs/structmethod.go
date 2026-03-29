package main

import "math"

type circle struct{
	radius int
}

// (c circle) - receiver
// Area - method name
// this is a method of the struct circle
// value receiver
func (c circle) Area() float64 {
	return math.Pi*float64(c.radius)*float64(c.radius)
}

func main(){

}