package main

import "github.com/kemal576/go-tcp/server"

func main() {
	s := server.NewServer()
	s.Listen()
}
