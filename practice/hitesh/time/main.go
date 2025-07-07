package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	myTime := time.Now()
	fmt.Println(myTime.Format(time.Stamp))
	fmt.Println(runtime.Version())
}
