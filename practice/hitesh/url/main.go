package main

import (
	"fmt"
	"net/url"
)

func main() {
	scheme := "https"
	domain := "example.com"
	path := "/path/to/resource"
	query := "foo=bar"

	url := url.URL{
		Scheme:   scheme,
		Host:     domain,
		Path:     path,
		RawQuery: query,
	}

	fmt.Println("My url", url.String())

}
