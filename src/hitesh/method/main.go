package main

import (
	"fmt"
)

func main() {
	hasib := &User{ //the & sing are user for pass the address of the variable
		FirstName: "Al",
		LastName:  "Hasib",
		Email:     "hasib@gmail.com",
		Age:       25,
		Verified:  true,
		Status:    "Active",
	}

	fmt.Println("The user Full name is:", hasib.GetFullName("Ahmed", "Shakil"))
	fmt.Printf("After modified %v %v \n", hasib.FirstName, hasib.LastName)
}

type User struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
	Verified  bool
	Status    string
}

func (u *User) GetFullName(fn, ln string) string { //the * use for pass value of the struct address
	u.FirstName = fn
	u.LastName = ln
	return u.FirstName + " " + u.LastName
}
