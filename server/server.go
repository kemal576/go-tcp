package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type Server struct {
	IP         net.IP
	Port       string
	Connection net.Conn
}

func NewServer() *Server {
	return &Server{
		IP:   findLocalIP(),
		Port: "5764",
	}
}

func (s *Server) Listen() {
	var adress string = s.IP.String() + ":" + s.Port
	ln, err := net.Listen("tcp", adress)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server started --> ", adress)

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("New Client -->", conn.RemoteAddr().String())
		go readInput(conn)
	}
}

func findLocalIP() net.IP {
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	addrs, err := net.LookupIP(host)
	if err != nil {
		panic(err)
	}

	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return ipv4
		}
	}
	return nil
}

func readInput(conn net.Conn) {
	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected -->", conn.RemoteAddr().String())
			break
		}

		msg = strings.Trim(msg, "\r\n")

		fmt.Println("New Message:", msg)
		fmt.Fprintf(conn, msg+"\n")
	}
}
