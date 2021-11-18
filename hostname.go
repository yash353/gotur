package main

import (
	"fmt"
	"os"
)

func main() {

	hostname, error := os.Hostname()

	if error != nil {

		panic(error)
	}

	fmt.Println("hostname returned from Environment : ", hostname)

	fmt.Println("error : ", error)

}
