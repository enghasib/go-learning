package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (user Person) printUser(){
	fmt.Printf("Namae: %s, Age: %d\n", user.Name, user.Age)
}

func main() {
	user:= Person{
		Name: "John",
		Age:  30,
	}
	user.printUser()

}