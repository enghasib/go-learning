package main

import "fmt"

func main() {
	Greeting()
	name := Name()
	sum, mul := Calc()
	fmt.Println("The sum result is", sum)
	fmt.Println("The mul result is", mul)
	Saygoodbye(name)
}
