package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Hitesh Conversion Program!")
	fmt.Println("Please enter the value you want to convert:")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	num, err := strconv.ParseFloat(strings.TrimSpace(input), 32)

	if err != nil {
		fmt.Println("Invalid input ", err)
	}

	// Trim the newline character from the input
	fmt.Println("Your modified value is:", num*num)

	fmt.Println("Thank you for using the Hitesh Conversion Program!")
}
