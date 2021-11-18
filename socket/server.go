package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	con, err := net.Listen("tcp", "localhost:8085")

	if err != nil {

		panic(err)
	}
	defer con.Close()

	for {

		c, err := con.Accept()

		if err != nil {

			panic(err)
		}
		fmt.Println("client connected")

		fmt.Println("client" + c.RemoteAddr().String() + "connected")

		go handleConnection(c)
	}
}

func handleConnection(conn net.Conn) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {

		fmt.Println("client left")
		conn.Close()
		return
	}

	log.Println("client message:", string(buffer[:len(buffer)-1]))

	conn.Write(buffer)

	handleConnection(conn)
}
