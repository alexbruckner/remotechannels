package remote

import (
	"fmt"
	"net"
	"bufio"
	"strconv"
)

func SendIntChan(address string) chan int {

	ch := make(chan int)

	go func(ch chan int) {
		for {
			select {
			case x := <-ch:
				send(strconv.Itoa(x), address)
			}
		}
	}(ch)

	return ch

}

func ReceiveStringChan(laddr string) chan string {
	ch := make(chan string)

	go func(ch chan string) {
		receive(ch, laddr)
	}(ch)

	return ch
}

func send (s string, address string) {
	// connect to this socket
	conn, err := net.Dial("tcp", address)

	if err != nil {
		panic (err)
	}

	fmt.Fprintf(conn, s+"\n")

	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(message)

}

func receive (ch chan string, laddr string) {
	ln, _ := net.Listen("tcp", laddr)

	for {
		conn, _ := ln.Accept()
		message, _ := bufio.NewReader(conn).ReadString('\n')
		ch <- message
		conn.Write([]byte("OK\n"))
		conn.Close()
	}
}

