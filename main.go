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

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	clientAddress := conn.RemoteAddr().String()
	fmt.Println("accepted connection from:", clientAddress)

	greeting := "welcome to the tcp server \r\n"
	conn.Write([]byte(greeting))

	fmt.Println("Closing connection from:", clientAddress)
}
