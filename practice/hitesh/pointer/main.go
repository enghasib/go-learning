package main

import (
	"fmt"
)

func main() {
	b := 255
	a := &b
	fmt.Println("address of b is", a)
	*a = 100

	fmt.Println("value of a is", *a)
	fmt.Println("The value of b ", b)
}
