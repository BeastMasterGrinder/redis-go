package main

import (
	"fmt"
	// Uncomment this block to pass the first stage
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
	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer c.Close()

	for {
		buf := make([]byte, 1024)

		//read message from the client
		r, err := c.Read(buf)
		if err != nil {
			fmt.Println("Error Reading from Connection")
			break
		}

		fmt.Printf("Read this from Client: %d", r)
		_, err = c.Write([]byte("+PONG\r\n"))

		if err != nil {
			fmt.Println("Error Writing to Connection")
			break
		}

	}

}
