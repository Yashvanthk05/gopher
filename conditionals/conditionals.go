package main

import "fmt"

func main() {
	age := 29

	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age < 25 {
		fmt.Println("You are a young adult.")
	} else {
		fmt.Println("You are an adult.")
	}

	//switch case

	switch age {
	case 18:
		fmt.Println("ur a minor")
		// break statement is implicit in go
		// if we want the case to fall through the next case then fallthrough keyword is used
	case 19:
		fallthrough
	case 20:
		fallthrough
	case 21:
		fmt.Println("U're in ur 20's")
	default:
		fmt.Print("ur none of the above")
	}
}
