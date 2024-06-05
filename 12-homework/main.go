package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn, id int) {
	defer conn.Close()
	fmt.Printf("Client %d connected.\n", id)
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Client %d disconnected.\n", id)
			return
		}
		fmt.Printf("Client %d says: %s", id, message)
		conn.Write([]byte(fmt.Sprintf("Received: %s", message)))
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer ln.Close()
	fmt.Println("Server started on :8080")

	clientCount := 0

	for clientCount < 3 {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		clientCount++
		go handleConnection(conn, clientCount)
	}
	fmt.Println("Reached 3 clients, no longer accepting new connections.")
}
