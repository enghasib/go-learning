package main

import "fmt"

func main() {
	day := []string{"Saturday", "Sunday", "Monday", "Tuesday", "WednesDay", "Thursday", "Friday"}

	// for i := 0; i < len(day); i++ { // traditional for loop
	// 	fmt.Printf("This is %v \n", day[i])
	// }

	// 	for i := range day { // for loop with range
	// 		fmt.Printf("This is %v \n", day[i])
	// 	}

	// for i, v := range day { // for loop with range and index
	// 	fmt.Printf("This is index %v and the value is %v \n", i, v)
	// }

	for i := 0; i <= len(day); i++ {
		if day[i] == "Friday" {
			fmt.Println("This is holy day")
			goto vacation
		}
	}
vacation:
	fmt.Println("I am going to vacation")
}
