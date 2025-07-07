package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak(sound string) string {
	return sound
}

func talk(a Animal, s string) {
	fmt.Println(a.Speak(s))
}

func main() {
	dog := Dog{}

}
