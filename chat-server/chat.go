package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go dM()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

var (
	messages = make(chan string)
)

func dM() {
	for {
		select {
		case msg := <-messages:
			fmt.Println(msg)
		}
	}
}

func handleConn(conn net.Conn) {
	who := conn.RemoteAddr().String()
	messages <- who + "has arrived"
	message, _ := bufio.NewReader(conn).ReadString('\n')
	messages <- message
	conn.Close()
}
