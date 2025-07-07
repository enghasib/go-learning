package main

import "fmt"

func main() {
	fmt.Println(calculator().Add(5, 5))
	fmt.Println(calculator().Sub(10, 5))
	fmt.Println(calculator().Mul(10, 5))
	fmt.Println(calculator().Div(10, 5))
}

type Calculate struct {
	Add func(int, int) int
	Sub func(int, int) int
	Mul func(int, int) int
	Div func(int, int) float32
}

func calculator() Calculate {
	return Calculate{
		Add: func(val1, val2 int) int {
			return val1 + val2
		},
		Sub: func(val1, val2 int) int {
			return val1 - val2
		},
		Mul: func(val1, val2 int) int {
			return val1 * val2
		},
		Div: func(val1, val2 int) float32 {
			if val2 == 0 {
				return float32(val1)
			} else {
				return float32(val1) / float32(val2)
			}
		},
	}

}
