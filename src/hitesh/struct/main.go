package main

import "fmt"

func main() {
	hasib := Person{
		Name:     "Hasib",
		Age:      25,
		Email:    "hasib@example.com",
		Verified: true,
	}

	fmt.Println(hasib)

}

type Person struct {
	Name     string
	Age      int
	Email    string
	Verified bool
}
