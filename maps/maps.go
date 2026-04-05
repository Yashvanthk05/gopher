package main

import "fmt"

func main(){
	// declaration
	stud := make(map[int]string)

	// adding keys
	stud[1] = "yashvanth"
	stud[2] = "deepak"
	fmt.Println(stud)

	for k,v := range stud {
		fmt.Println(k,v)
	}

	// deleting keys
	delete(stud,1)

	// checking existence
	if _, ok := stud[1]; !ok {
		fmt.Println("Student with ID 1 does not exist")
	}else {
		fmt.Println("Student with ID 1 exist")
	}

	// nil map - declared but not initialized
	var db map[string]int
	db["yash"] = 1 // this will panic since no underlying buckets, safe method is to initialize using make
	// map is a reference type it stores reference (pointer) to the data structure no actual data is stored
	m1 := map[string]int{
		"a": 1,
	}

	m2 := m1
	m2["a"] = 100

	fmt.Println(m1["a"]) // 100 since reference is changed becoz m1 and m2 both point to the same underlying map
}