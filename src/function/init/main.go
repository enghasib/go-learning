package main

import "fmt"

var a int = 10;

func main(){
	fmt.Println(a)
}

func init (){
	fmt.Println(a);

	a = 20
}