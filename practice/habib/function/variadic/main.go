package main

import "fmt"

func sum(nums ...int) int {
	res := 0
	for _, value := range nums {
		res += value
	}
	return res
}

func main() {
	fmt.Println(sum(1, 3, 4, 5, 6))
}
