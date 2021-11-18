package main

import (
	"fmt"
)

func main() {

	fmt.Println("main start")

	c := make(chan string)

	go ch(c)

	c <- "yash"

	fmt.Println("main stop")

}

func ch(c chan string) {

	fmt.Println("Hello " + <-c + " !!!")

}
