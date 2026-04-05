package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Person struct {
	Name     string `json:"name"`
	Age      int    `json:"age,omitempty"`
	Password string `json:"-"`
	Email    string `json:"email"`
}

func main() {
	p1 := Person{Name: "Yashvanth", Age: 21, Password: "1234", Email: "dev@gmail.com"}
	// json.Marshal - Go data structure to JSON object
	data, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(errors.New("Error in JSON encoding"))
		return
	}
	fmt.Println(string(data))

	data = []byte(`{"name":"Deepak","age":18,"email":"deepu@gmail.com"}`)
	var p2 Person
	// json.Unmarshal - JSON object to Go data Structure; passes a pointer so that data can be stored in place
	json.Unmarshal(data,&p2)
	fmt.Println(p2)

	// instead of loading everything into memory, directly stream to/from a writer(file, HTTP)
	// json.Encoder(w).Encode(p1)
	// json.Decoder(r.Body).Decode(&p1)


	// problem with non-pointer fields: If you unmarshal {"age": 0} and {} (age absent), both give you p.Age == 0. You can't tell the difference

	/*
	type Profile struct {
		Age *int `json:"age"`
	}

	{"age": 30} *p.Age == 30 Field present, value 30 
	{"age": 0} *p.Age == 0 Field present, value 0
	{"age": null} p.Age == nil Field explicitly null
	{} p.Age == nil Field absent entirely
	*/


}
