package main

import (
	"./bru/remote"
	"fmt"
)

func main() {

	var ch chan string = remote.ReceiveStringChan(":8080")
	for {
		s := <- ch
		fmt.Printf("ReceiveStringChan: %v", s)
	}
}

