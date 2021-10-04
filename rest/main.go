package main

import "time"

//this starts a server in a separate goroutine and then starts a client that sends request to the server.
//output via terminal.
func main() {
	go server()
	time.Sleep(2 * time.Second)
	client()
}
