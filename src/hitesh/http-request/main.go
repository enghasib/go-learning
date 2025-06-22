package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Simple get request using http package in go")
	// CreateGetReq()
	// CreatePostReq()
	PerformPostJsonRequest()
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

func PerformPostJsonRequest() {
	myUrl := "http://localhost:5000/postForm"

	data := url.Values{}
	data.Add("firstName", "Al")
	data.Add("lastName", "Hasib")
	data.Add("age", "25")
	data.Add("email", "hasib@gmail.com")

	//make a request
	res, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("The response body is: ", string(content))
}
