package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8085")

	if err != nil {

		panic(err)
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Text to send: ")

		input, _ := reader.ReadString('\n')

		fmt.Println(input)

		message, _ := bufio.NewReader(conn).ReadString('\n')

		log.Print("Server relay:", message)
	}
}
