package main

import "fmt"

func addName(db map[rune]map[string]int, name string) {
	firstchar := rune(name[0])
	_, ok := db[firstchar]
	if !ok {
		db[firstchar] = make(map[string]int)
	}
	db[firstchar][name]++
}

func main(){
	db := make(map[rune]map[string]int)
	addName(db,"Yashvanth")
	addName(db,"Yashix")
	fmt.Println(db)
}