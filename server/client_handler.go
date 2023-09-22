package server

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"sync"
)

// Message represents the structure of the message sent by clients.
type Message struct {
	Text string `json:"text"`
}

// Acknowledgment represents the acknowledgment message sent by the server.
type Acknowledgment struct {
	Message string `json:"message"`
}

// HandleClient handles an individual client connection.
func HandleClient(conn net.Conn, mutex *sync.Mutex) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				log.Printf("Client disconnected\n")
			} else {
				log.Printf("Error reading from client: %v\n", err)
			}
			return
		}

		var clientMessage Message
		if err := json.Unmarshal(buffer[:n], &clientMessage); err != nil {
			log.Printf("error decoding JSON: %v\n", err)
			continue
		}

		log.Printf("Received message from client: %s\n", clientMessage.Text)

		ack := Acknowledgment{Message: "Message received successfully"}

		ackJSON, err := json.Marshal(ack)
		if err != nil {
			log.Printf("error encoding JSON: %v\n", err)
			continue
		}

		mutex.Lock()
		_, err = conn.Write(ackJSON)
		mutex.Unlock()

		if err != nil {
			log.Printf("error sending acknowledgment: %v\n", err)
			return
		}
	}
}
