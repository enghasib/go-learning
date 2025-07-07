package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	// var myList = []int{}
	// myList = append(myList, 1, 2, 3, 4)
	// fmt.Println(myList)

	// fruitList := []string{"apple", "banana", "cherry"}
	// fruitList = append(fruitList[1:])

	// fmt.Println(fruitList)

	// nameList := make([]string, 0)
	// nameList = append(nameList, "hasib", "shakil", "mamun")

	// fmt.Println(nameList)

	sortList := make([]int, 0)

	sortList = append(sortList, 324)
	sortList = append(sortList, 328)
	sortList = append(sortList, 428)
	sortList = append(sortList, 228)
	sortList = append(sortList, 528)
	sortList = append(sortList, 28)

	areSorted := sort.IntsAreSorted(sortList)
	fmt.Println("Is slice are sorted:", areSorted)

	slices.Sort(sortList)
	fmt.Println("Sort slice:", sortList)

}
