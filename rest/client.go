package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Now demoing a GET request")
	fmt.Println(getDemo())
	fmt.Println("---------------")
	time.Sleep(1 * time.Second)
	fmt.Println("Now demoing a POST request")
	fmt.Println(postDemo())
	fmt.Println("---------------")
	time.Sleep(1 * time.Second)
	fmt.Println("Now demoing a DELETE request")
	fmt.Println(deleteDemo())
	fmt.Println("---------------")
	time.Sleep(1 * time.Second)
	fmt.Println("Now demoing a PUT request")
	fmt.Println(putDemo())

}

func getDemo() string {
	resp, err := http.Get("http://localhost:8080/courses")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	bodyString := string(body)
	return bodyString

}

func postDemo() string {
	client := &http.Client{}
	var jsonStr = []byte(`{"id": "5", "rating": 0.1, "name": "Bæredygtigt design", "workload": 0.1, "students": []}`)
	//make request
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/courses", bytes.NewBuffer(jsonStr)) //send the json string as payload
	req.Header.Set("X-Custom-Header", "myvalue")                                                            //wtf
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
	}
	//fetch request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer resp.Body.Close()
	//read response
	body, err := io.ReadAll(resp.Body)
	bodyString := string(body)
	return bodyString
}

func deleteDemo() string {
	client := &http.Client{}
	//make request
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/courses/1", nil)
	if err != nil {
		fmt.Println(err)
	}
	//fetch request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer resp.Body.Close()
	//read response
	body, err := io.ReadAll(resp.Body)
	bodyString := string(body)
	return bodyString
}

//does not output anything atm
func putDemo() string {
	client := &http.Client{}
	var jsonStr = []byte(`{"id": "4", "rating": 9.5, "name": "Algo", "workload":7.5, "students": ["Alle", "Er med"]}`)
	//make request
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8080/courses/2", bytes.NewBuffer(jsonStr)) //send the json string as payload
	//req.Header.Set("X-Custom-Header", "myvalue") //wtf
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
	}
	//fetch request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer resp.Body.Close()
	//read response
	body, err := io.ReadAll(resp.Body)
	bodyString := string(body)
	return bodyString
}
