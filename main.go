package main

import (
	"goServer/server"
	"log"
)

func main() {
	const (
		HOST = "127.0.0.1"
		PORT = "8080"
	)
	if err := server.StartServer(HOST, PORT); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
