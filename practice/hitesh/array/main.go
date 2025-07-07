package main

import "fmt"

func main() {
	// var myArray [4]string
	// myArray[0] = "Hello"
	// myArray[1] = "World"
	// myArray[2] = "This"
	// myArray[3] = "Is"

	var myArray []int
	myArray = append(myArray, 1)
	myArray = append(myArray, 2)
	myArray = append(myArray, 3)
	myArray = append(myArray, 4)

	fmt.Println("This is the array: ", myArray)

}
