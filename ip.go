package main

import (
	"fmt"
	"net"
)

func main() {

	conn, _ := net.Dial("udp", "8.8.8.8:")

	defer conn.Close()

	ipAddress := conn.LocalAddr().(*net.UDPAddr)

	fmt.Println("Machine ip address :", ipAddress)
}
