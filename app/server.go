package main

import (
	"fmt"
	// Uncomment this block to pass the first stage
	"io"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	for {

		c, err := l.Accept()
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}

		go handleRequest(c)
		defer c.Close()
	}

}

func handleRequest(c net.Conn) {
	command := make([]byte, 1024)


	r, err := c.Read(command)

	if err != nil {
		fmt.Println("Error Reading from connection")
		os.Exit(1)
	}

	fmt.Printf("Read: \n%s", string(r))
	_, err = c.Write([]byte("+PONG\r\n"))

	if err != nil {
		fmt.Println("Error sending response")
		os.Exit(1)
	}
}
