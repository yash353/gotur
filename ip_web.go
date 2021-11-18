package main

import (
	"fmt"
	"net"
)

func main() {

	addresh, err := net.LookupIP("youtube.com")

	if err != nil {

		fmt.Println("unknow error occur :", err)
	} else {
		fmt.Println("addresh of site is  :", addresh)
	}
}
