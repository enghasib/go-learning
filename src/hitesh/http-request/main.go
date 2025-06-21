package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Simple get request using http package in go")
	// CreateGetReq()
	CreatePostReq()
}

func CreateGetReq() {
	url := "http://localhost:5000/get"

	var stringBuilder strings.Builder

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Res status code:", response.Status)
	fmt.Println("Res status code:", response.ContentLength)

	content, _ := io.ReadAll(response.Body)
	// fmt.Println(string(redRes))

	stringBuilder.Write(content)

	fmt.Println(stringBuilder.String())

	defer response.Body.Close()
}

func CreatePostReq() {
	url := "http://localhost:5000/post"

	// request body payload
	requestBody := strings.NewReader(`
	{
		"name":"hasib",
		"email": "hasib@gmail.com",
		"age": 25
	}
	`)

	// make a post request
	res, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	content, _ := io.ReadAll(res.Body)

	fmt.Println("The response is: ", string(content))

}
