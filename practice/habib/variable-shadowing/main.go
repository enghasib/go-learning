package main

import "fmt"

var a = 10

func main() {
	age := 20
	if age > 18 {
		a := 50
		fmt.Println("local scope:", a)
	}
	fmt.Println("Global scope:", a)
}
