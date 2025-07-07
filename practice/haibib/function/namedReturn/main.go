package main

import "fmt"

func calculator(a, b int) (mul, div int) {
	mul = a * b
	div = a / b
	return mul, div
}

func main() {
	mul, div := calculator(10, 2)
	fmt.Println("Multiplication", mul)
	fmt.Println("Division", div)
}
