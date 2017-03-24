package main

import (
	"time"
	"./bru/remote"
	"fmt"
)

func main() {

	var ch chan int = remote.SendIntChan("test", "127.0.0.1:8080")
	var ch2 chan int = remote.SendIntChan("test2", "127.0.0.1:8080")

	for i := 1; i <= 10; i++ {
		fmt.Printf("SendIntChan: %v\n", i)
		ch <- i
		ch2 <- i
		time.Sleep(1 * time.Second)
	}

}
