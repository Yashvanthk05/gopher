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
	employee   // embedding the employee struct - achievees composition (reusing behavior) instead of inheritance
	department string
}

type person struct {
	name string
	age  int
}

type earth struct {
	people []person
}

// struct tags - metadata annotations attached to struct fields (used by libraries such as JSON, ORM, validators)

type Person struct {
	Name 		string	`json:"name" db:"name"`
	Age 		int 	`json:"age,omitempty"`
	Password	string	`json:"-" db:"password"`	 
}

// using tags {"name":"John","age":25}
// without tags {"Name":"John","Age":25}
// rename field - Name string `json:"full_name"`
// Omit empty values - Age int `json:"age,omitempty"`
// ignore field - Password string `json:"-"`

func main() {
	// struct are value type
	stud1 := student{name: "yashvanth", rollno: 1589}
	fmt.Println(stud1)
	// creating an instance of the nessted college struct and initializing it with values
	college1 := college{}
	college1.name = "ABC College"
	college1.location = "New York"
	college1.student.name = "yashvanth"
	college1.student.rollno = 1589

	// positional struct literal and pointer to struct
	stud2 := &student{"kali",1320}
	// another way to create pointer
	stud3 := new(student)
	
	fmt.Println(stud3)

	// go automatically dereferences pointer to struct
	fmt.Println(stud2.name)

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

	manager2 := manager{
		employee: employee{
			name: "yashvanth",
			age: 21,
		},
		department: "scope",
	}

	fmt.Println(manager2)
	// Go approach (composition)
	// manager has a employee

	// creating an instance of the earth struct and initializing it with a slice of person structs
	earth1 := earth{}
	// named struct literal
	earth1.people = []person{
		{name: "Alice", age: 30},
		{name: "Bob", age: 25},
		{name: "Charlie", age: 35},
	}
	fmt.Println(earth1.people[0].name)
}
