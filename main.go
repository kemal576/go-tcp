package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":5764")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server started --> localhost:5764")

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("New Client -->", conn.RemoteAddr().String())
		go readInput(conn)
	}
}

func readInput(conn net.Conn) {
	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected -->", conn.RemoteAddr().String())
			break
		}

		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/nick":
			fmt.Printf("New Command: /nick %s\n", args[1])
			fmt.Fprintf(conn, msg+"\n")

		case "/join":
			fmt.Printf("New Command: /join %s\n", args[1])
			fmt.Fprintf(conn, msg+"\n")

		default:
			fmt.Printf("Unknown Command: %s\n", cmd)
			fmt.Fprintf(conn, msg+"\n")
		}
	}
}
