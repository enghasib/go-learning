package main

import "fmt"

func namedReturn() (result int) {
	fmt.Println("Initial result:", result)
	show := func() {
		result = result + 10
		fmt.Println("Defer result:", result)
	}
	defer show()
	result = 5
	return
}

func typeReturn() int {
	result := 0
	fmt.Println("Initial result:", result)
	show := func() {
		result = result + 10
		fmt.Println("Defer result:", result)
	}
	defer show()
	result = 5
	return result
}

func main() {
	namedReturn := namedReturn()
	fmt.Println("Main call named:", namedReturn)

	typeReturn := typeReturn()
	fmt.Println("Main call typed:", typeReturn)
}
