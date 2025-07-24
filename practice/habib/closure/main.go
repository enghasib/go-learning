package main

import "fmt"

const a int = 10

var p int = 20

func outer() func() {
	var money int = 100
	var age int = 30
	fmt.Println("Age is:", age)

	show := func() {
		money = money + a + p
		fmt.Println(money)
	}
	return show
}

func call() {
	incr1 := outer()
	incr1()
	incr1()
	incr1()
}

func main() {
	call()
}
