package main

import (
	"encoding/json"
	"goServer/server"
	"net"
	"testing"
	"time" 
	"errors" 
)

func TestServer(t *testing.T) {
	// Start the server in a goroutine.
	host := "127.0.0.1"
	port := "8080"
	go func() {
		if err := server.StartServer(host, port); err != nil {
			t.Errorf("Server error: %v", err)
			return
		}
	}()

	// Wait for the server to start listening.
	conn, err := waitForServer(host, port)
	if err != nil {
		t.Fatalf("Server not reachable: %v", err)
	}
	defer conn.Close()

	t.Run("TestSendMessage", func(t *testing.T) {
		// Create a test message.
		testMessage := server.Message{Text: "Hey i am server test"}

		// Send the test message to the server.
		sendMessage(conn, testMessage)

		// Receive and validate the acknowledgment message from the server.
		acknowledgment := receiveMessage(conn)
		if acknowledgment.Message != "Message received successfully" {
			t.Errorf("Expected acknowledgment message 'Message received successfully', got '%s'", acknowledgment.Message)
		}
	})
}

func waitForServer(host, port string) (net.Conn, error) {
	// Implement a loop with retries to wait for the server to start.
	for i := 0; i < 10; i++ { // Retry for a maximum of 10 seconds.
		conn, err := net.Dial("tcp", host+":"+port)
		if err == nil {
			return conn, nil
		}
		time.Sleep(1 * time.Second) // Wait for 1 second before retrying.
	}
	return nil, errors.New("server not reachable")
}

func sendMessage(conn net.Conn, msg server.Message) {
	msgJSON, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	_, err = conn.Write(msgJSON)
	if err != nil {
		panic(err)
	}
}

func receiveMessage(conn net.Conn) server.Acknowledgment {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}

	var acknowledgment server.Acknowledgment
	err = json.Unmarshal(buffer[:n], &acknowledgment)
	if err != nil {
		panic(err)
	}

	return acknowledgment
}
