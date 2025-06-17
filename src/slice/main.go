package main

import "fmt"

func main(){
//That's the example of a array and loop
// numArr := [5]string{"hasib","shakil","mamun","torik","golam"} 

// for  value:= range numArr {
//     fmt.Println(numArr[value])
// }

// slice
    s:= []int{}
s = append(s, 1)
s = append(s, 2)
s = append(s, 3)
s = append(s, 5)

println("this is the range of the slice: ",len(s))


fmt.Println(s)
}

                     