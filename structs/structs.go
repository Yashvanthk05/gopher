package main

import "fmt"

type student struct {
	name   string
	rollno int
}

// nested struct
type college struct {
	name     string
	location string
	student  student
}

// embedded struct
type employee struct {
	name string
	age  int
}

type manager struct {
	employee   // embedding the employee struct into the manager struct
	department string
}

type person struct {
	name string
	age  int
}

type earth struct {
	people []person
}

func main() {
	stud1 := student{name: "yashvanth", rollno: 1589}
	fmt.Println(stud1)
	// creating an instance of the nessted college struct and initializing it with values
	college1 := college{}
	college1.name = "ABC College"
	college1.location = "New York"
	college1.student.name = "yashvanth"
	college1.student.rollno = 1589

	stud2 := student{
		name:   "kali",
		rollno: 1320,
	}

	fmt.Println(stud2)

	col2 := college{
		name:     "vit",
		location: "vandaloor",
		student: student{
			name:   "kavi",
			rollno: 16,
		},
	}

	fmt.Println(col2)
	

	fmt.Println(college1)
	// anonymous struct
	// an anonymous struct is a struct that does not have a name and is defined inline. It is useful when you need to create a one-time struct for a specific purpose without the need to define a separate type.
	// only one instance of the anonymous struct can be created and it cannot be reused.
	car := struct {
		make  string
		model string
		year  int
	}{
		make:  "Toyota",
		model: "Camry",
		year:  2020,
	}
	fmt.Println(car)

	// creating an instance of embedded struct
	manager1 := manager{}
	manager1.name = "John Doe" // no need for manager1.employee.name because the employee struct is embedded in the manager struct, so we can access its fields directly from the manager struct
	manager1.age = 35
	manager1.department = "Sales"
	fmt.Println(manager1)

	// creating an instance of the earth struct and initializing it with a slice of person structs
	earth1 := earth{}
	earth1.people = []person{
		{name: "Alice", age: 30},
		{name: "Bob", age: 25},
		{name: "Charlie", age: 35},
	}
	fmt.Println(earth1)
}
