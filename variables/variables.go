package main

import "fmt"

func main() {
	var name string = "Yashvanth"
	var age int = 21
	const pi float64 = 3.14159
	dummy := "This is a dummy variable." // short assignment operator
	
	fmt.Println("Hello, World!")
	fmt.Printf("The value of pi is approximately %.5f.\n", pi)
	fmt.Println(dummy)
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// types used in resource allocation and management
	var a int8 = 127
	var b int16 = 32767
	var c int32 = 2147483647
	var d int64 = 9223372036854775807
	var e uint8 = 255
	var f uint16 = 65535
	var g uint32 = 4294967295
	var h uint64 = 18446744073709551615
	
	// default int type is either 32 or 64 bits depending on the platform
	var i complex128 = 1 + 2i
	var j bool = true
	var k byte = 'A' // byte is an alias for uint8
	var l rune = '世' // rune is an alias for int32, used for Unicode code points

	fmt.Printf("Integer types: %d, %d, %d, %d\n", a, b, c, d)
	fmt.Printf("Unsigned integer types: %d, %d, %d, %d\n", e, f, g, h)
	fmt.Printf("Complex number: %v\n", i)
	fmt.Printf("Boolean value: %t\n", j)
	fmt.Printf("Byte value: %c\n", k)
	fmt.Printf("Rune value: %c\n", l)

	msg := fmt.Sprintf("My name is %s and I am %d years old.", name, age)
	fmt.Println(msg)

	x := 10
	y := x
	// in go variables are passed by value not by reference, so y is a copy of x
	fmt.Printf("x: %d, y: %d\n", x, y)
	x = 20
	fmt.Printf("After changing x: x: %d, y: %d\n", x, y)

	// semi-colons are optional in Go, they are automatically inserted at the end of lines
}