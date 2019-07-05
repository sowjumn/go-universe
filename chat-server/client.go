package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	myClient()
}

func myClient() {
	conn, _ := net.Dial("tcp", "localhost:8000")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Hi: ")
	for scanner.Scan() {
		fmt.Fprintf(conn, scanner.Text()+"\n")
	}

	conn.Close()
}
