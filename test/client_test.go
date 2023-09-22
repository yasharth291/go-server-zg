package main

import (
	"goServer/server"
	"net"
	"sync"
	"testing"
)

func TestClientCommunication(t *testing.T) {
	// Start a test server in a goroutine.
	host := "127.0.0.1"
	port := "8081"
	go func() {
		if err := startTestServer(host, port); err != nil {
			t.Errorf("Test server error: %v", err)
			return
		}
	}()

	// Wait for the test server to start listening.
	conn, err := waitForServer(host, port)
	if err != nil {
		t.Fatalf("Test server not reachable: %v", err)
	}
	defer conn.Close()

	// Create a test message.
	testMessage := server.Message{Text: "hey i am test"}

	// Send the test message to the test server.
	sendMessage(conn, testMessage)

	// Receive and validate the acknowledgment message from the test server.
	acknowledgment := receiveMessage(conn)
	if acknowledgment.Message != "Message received successfully" {
		t.Errorf("expected acknowledgment message 'Message received successfully', got '%s'", acknowledgment.Message)
	}
}

func startTestServer(host, port string) error {
	listenAddr := host + ":" + port

	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	defer listener.Close()

	var mutex sync.Mutex

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go server.HandleClient(conn, &mutex)
	}
}
