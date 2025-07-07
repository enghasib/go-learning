package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch) // close when done sending
	}()

	for val := range ch {
		fmt.Println(val)
	}

}
