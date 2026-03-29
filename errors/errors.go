package main

import (
	"fmt"
	"errors"
)

// error is a built in interface with Error() method that returns string

type customerr struct{
	msg string
}

func (c customerr) Error() string {
	return fmt.Sprintf(c.msg)
}

func divide(a, b int) (float64,error) {
	// custom error handler
	if b == 0 {
		var c error = customerr{fmt.Sprintf("Cannot divide %v by zero",a)}
		return 0.0, c
	}
	if b==0 {
		return 0.0, errors.New("Cannot divide by Zero!")
	}
	return float64(a)/float64(b),nil
}

func main(){
	c,err := divide(3,0)
	if err!=nil {
		fmt.Println("Error:",err)
		return
	}
	fmt.Println("Result:",c)
}