package main

import "fmt"

func worker(ch chan string) {
	name := <-ch // receiving value from channel
	fmt.Println("Hello", name)
}

func calculate(ch chan float64) {
	cal := <-ch
	cal = cal * cal
	fmt.Println(cal)
}

func main() {
	ch := make(chan string)
	calCh := make(chan float64)
	go worker(ch) // launch goroutine
	ch <- "Hasib" // send data via channel

	go calculate(calCh)
	calCh <- 5.0
}
