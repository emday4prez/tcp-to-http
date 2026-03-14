package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//open tcp socket
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Failed to bind to port 8080:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("raw stream inspector listening on 127.0.0.1:8080...")

	//wait for single client to connect
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("error accepting connection:", err)
		os.Exit(1)
	}
	// valid HTTP response requires a Status Line, Headers, a blank line, and a Body.
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"Content-Length: 21\r\n" +
		"Connection: close\r\n" +
		"\r\n" + // separates headers from the body
		"Hello from raw bytes!"

	conn.Write([]byte(response))

	defer conn.Close()
	fmt.Println("Client connected from:", conn.RemoteAddr())

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from stream:", err)
		return
	}
	fmt.Println("--- RAW BYTES RECEIVED ---")
	fmt.Print(string(buffer[:n]))
	fmt.Println("--------------------------")
}
