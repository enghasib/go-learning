package main

import "fmt"

func Greeting() {
	fmt.Println("Hey Welcome to my new application!")
}

func Name() string {
	var name string
	fmt.Println("Enter your name:-")
	fmt.Scanln(&name)
	return name
}

func Calc() (sum int, mul int) {
	var (
		num1 int
		num2 int
	)
	fmt.Println("Enter the first number:")
	fmt.Scanln(&num1)

	fmt.Println("Enter the second number:")
	fmt.Scanln(&num2)

	sumResult := num1 + num2
	mulResult := num1 * num2

	return sumResult, mulResult
}

func Saygoodbye(name string) {
	fmt.Printf("Thank you %s for using my application. Goodbye!", name)
}
