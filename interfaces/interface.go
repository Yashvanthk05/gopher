package main

import ("fmt";"math")

// interfaces are collection of method signatures
// interface - collection of method definitions (no implementation)

// achieve polymorphism by using interface as a function parameter
// polymorphism means one function can work with multiple types
// same function for multiple types - less duplication
// decoupling - code that depends on behavior, not concrete types

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

func (c Circle) Area() float64{
	return math.Pi*c.radius*c.radius
}

func (r Rect) Area() float64{
	return r.width*r.height
}

// interfaces are implicitly implemented
// one function works for many type
func printArea(s Shape){
	fmt.Println(s.Area())
}

// multiple interfaces

type Scanner struct{
	pages int
}

type Reader interface{
	Scan()
}

type Writer interface{
	Print()
}

func (s Scanner) Scan(){
	fmt.Println("Scanning:",s.pages)
}

func (p Scanner) Print(){
	fmt.Println("Reading:",p.pages)
}

// combined interface
type ReaderWriter interface{
	Print()
	Scan()
}

// empty interface - accepts value of any type
func printValue(v interface{}){
	fmt.Println(v)
}

func main(){
	c1 := Circle{10}
	printArea(c1)
	r1 := Rect{4,5}
	printArea(r1)

	// multiple instances
	var s Reader = Scanner{10}
	s.Scan()
	var p Writer = Scanner{20}
	p.Print()
	var ps ReaderWriter = Scanner{30}
	ps.Print()
	ps.Scan()

	// type assertion - get actual value (and type) stored inside an interface
	// interface can hold any type
	// type assertion allows to extract the real type from it
	// val := i.(Type); i->Interface variable, Type->Actual type we expect
	// Type Assertion = “Open the box and check”
	// s := i.(string)
	// “I think this box has a string — give it to me as a string”
	// if our guess is wrong program will crash(panic)
	c, ok := s.(Writer)
	if ok {
		fmt.Println("It is a Scanner Type:",c)
	}else {
		fmt.Println("It is not a Scanner Type")
	}
	printValue("yashvanth")
	printValue(10)
	printValue(10.23232)
}