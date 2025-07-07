package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "This is the content that will be written to the file."

	file, err := os.Create("./myFile.txt")
	checkError(err)

	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Length of the content: %d\n", length)

	defer file.Close()
	readFile("./myFile.txt")
}

func readFile(file string) {
	content, err := os.ReadFile(file)
	checkError(err)

	fmt.Println("Read file content: ", string(content))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
