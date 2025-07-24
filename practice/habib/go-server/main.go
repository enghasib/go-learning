package main

import (
	"fmt"
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	reqBody, err := io.ReadAll(body)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(reqBody)
}

func main() {
	app := http.ser

}
