package server

import (
	"log"
	"net"
	"sync"
)

// StartServer starts the Unix domain socket server.
func StartServer(host, port string) error {
	listenAddr := host + ":" + port
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Printf("error listening on %s: %v", listenAddr, err)
		return err
	}
	defer listener.Close()

	log.Printf("Server listening on %s\n", listenAddr)

	var mutex sync.Mutex

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error accepting client connection: %v\n", err)
			continue
		}

		go func(conn net.Conn) {
			defer conn.Close()
			HandleClient(conn, &mutex)
		}(conn)
	}
}
