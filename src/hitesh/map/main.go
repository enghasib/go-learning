package main

import "fmt"

func main() {
	language := make(map[string]string)

	language["JS"] = "JavaScript"
	language["PY"] = "Python"
	language["CPP"] = "C++"
	language["PHP"] = "PHP"

	_, ok := language["js"]

	if ok {
		fmt.Println("The are is exist in the map")
	} else {
		fmt.Println("The key are not exist in the map")
	}

	// if i != "" {
	// 	fmt.Println("The key are exist in the map! ok is ->", ok)
	// } else {
	// 	fmt.Println("The key are not exist in the map! ok is ->", ok)
	// }
}
