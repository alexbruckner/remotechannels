package main

import (
	"./bru/remote"
	"fmt"
	"time"
)

func main() {

	go func() {
		var ch chan string = remote.ReceiveStringChan("test", ":8080")
		for {
			s := <-ch
			fmt.Printf("ReceiveStringChan1: %v\n", s)
		}
	}()

	go func() {
		var ch chan string = remote.ReceiveStringChan("test2", ":8080")
		for {
			s := <-ch
			fmt.Printf("ReceiveStringChan2: %v\n", s)
		}
	}()

	time.Sleep(60*time.Second)

}

