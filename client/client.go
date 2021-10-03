package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {
	resp, err := http.Get("http://localhost:8080/courses")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close() 
	body, err := io.ReadAll(resp.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}