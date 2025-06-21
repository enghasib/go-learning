package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://example.com"

func main() {
	//make a get request with the url
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	resByte, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("The response is: ", string(resByte))

	defer resp.Body.Close()
}
