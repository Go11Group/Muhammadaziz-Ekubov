package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to server. Type a message and press Enter:")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Server response: " + message)
	}
}
