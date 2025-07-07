package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Email    string   `json:"email"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("This is the lesson of json!")
	DecodeJson()
}

func EncodeJson() {
	persons := []person{{Name: "hasib", Age: 25, Email: "hasib@gmail.com", Password: "pass123", Tags: []string{"web-developer", "backend-engineer"}}}

	persons = append(persons, person{Name: "Shakil", Age: 24, Email: "shakil@gmail.com", Password: "pass456", Tags: nil})
	persons = append(persons, person{Name: "Mamun", Age: 24, Email: "mamun@gmail.com", Password: "pass789", Tags: []string{"muhaddes", "lecturer", "teacher"}})

	myJson, err := json.MarshalIndent(persons, "", " ")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", myJson)
}

func DecodeJson() {
	jsonData := []byte(`
		 {
			"name": "hasib",
			"age": 25,
			"email": "hasib@gmail.com",
			"tags": ["web-developer","backend-engineer"]
		}
	`)

	var newPerson person

	// check if the data is valid json
	isJsonValid := json.Valid(jsonData)
	if !isJsonValid {
		panic("Jason is not valid")
	}

	json.Unmarshal(jsonData, &newPerson)

	fmt.Printf("%v#\n", newPerson)

}
