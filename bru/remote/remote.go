package remote

import (
	"fmt"
	"net"
	"bufio"
	"strconv"
	"encoding/json"
)

func SendIntChan(name, address string) chan int {

	ch := make(chan int)

	go func(ch chan int) {
		for {
			select {
			case x := <-ch:
				send(name, address, strconv.Itoa(x))
			}
		}
	}(ch)

	return ch

}

func ReceiveStringChan(name, laddr string) chan string {
	ch := make(chan string)

	go func(ch chan string) {
		receive(ch, name, laddr)
	}(ch)

	return ch
}

type Message struct {
	Name string `json: name`
	Cargo string `json: cargo`
}


func send (name , address, s string) {
	// connect to this socket
	conn, err := net.Dial("tcp", address)

	if err != nil {
		panic (err)
	}

	message := new(Message)
	message.Name = name
	message.Cargo = s
	b, err := json.Marshal(message)

	if err != nil {
		panic (err)
	}

	fmt.Fprintf(conn, string(b)+"\n")

	status, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)

}

var receivers = make(map[string]bool)
var chans = make(map[string]chan string)

func receive (ch chan string, name, laddr string) {

	if chans[laddr+name] == nil {

		chans[laddr+name] = ch

		if !receivers[laddr] {
			receivers[laddr] = true
			go listen(laddr)
		}
	} else {
		fmt.Errorf("There is already a channel named [%v] running on [%v]", name, laddr)
	}
}

func listen(laddr string) {
	ln, _ := net.Listen("tcp", laddr)
	for {
		conn, _ := ln.Accept()
		message, _ := bufio.NewReader(conn).ReadString('\n')

		m := new(Message)
		json.Unmarshal([]byte(message), &m)

		chans[laddr+m.Name] <- m.Cargo
		conn.Write([]byte("OK\n"))
		conn.Close()
	}
}

