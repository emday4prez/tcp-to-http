package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
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
	fmt.Println("New connection from:", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading from client", err)
			return
		}

		fmt.Print("read line:", line)

		if strings.TrimSpace(line) == "" {
			fmt.Println("---- end of http headers ----")
			break
		}
	}

	response := "HTTP/1.1 200 OK\r\nContent-Length: 13\r\nConnection: close\r\n\r\nHello, World!"
	conn.Write([]byte(response))

}
