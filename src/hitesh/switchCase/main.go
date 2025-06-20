package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//generate random number between 1 and 6
	randomNumber := rand.Intn(6) + 1
	fmt.Println("The dice number is :", randomNumber)
	switch randomNumber {
	case 1:
		fmt.Println("Congrats, You can Start moving!")
		fallthrough
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
	case 4:
		fmt.Println("You can move 4 spot")
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("Well done! You get another chance to roll the dice.")
	default:
		fmt.Println("What is this ? Dice again!")
	}
}
