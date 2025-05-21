package main

import (
	"fmt"
	"time"
)

func printValue(number *[5] int){
    fmt.Println("The array is:", number)
}

func main(){
    var numberArray = [5] int {1,2,3,4,5};
    start:= time.Now()
    printValue(&numberArray)
    end := time.Now()
    
    delay := end.Sub(start)

    fmt.Printf("Time elapsed: %v\n", delay)
}