package main

import "fmt"

// interface value = (type, value) pair
// stores two things (interface) = (concrete type, concrete value)

// var x interface{} = 10
// x = (int, 10)


func main(){
	// case 1: type matters
	var x interface{} = 10
	var y interface{} = 20
	fmt.Println(x==y)

	// case 2: same value, different type
	var a interface{} = 10
	var b interface{} = int32(10)
	fmt.Println(a==b)

	// Nil interface vs interface holding nil pointer
	// nil interface
	var p1 interface{}
	fmt.Println(p1) // internally p1 = (type: nil, value: nil)
	fmt.Println(p1==nil)

	// interface holding nil pointer
	var p2 *int
	var p3 interface{} = p2
	fmt.Println(p3) // internally p3 = (type: *int, value: nil) interface is not fully nil only the value part is nil
	fmt.Println(p3==nil)
}