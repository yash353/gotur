package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func main() {

	s := bufio.NewScanner(os.Stdin)

	fmt.Println("enter no..")

	s.Scan()

	input, _ := strconv.ParseInt(s.Text(), 10, 64)

	fmt.Println(input)

	// fmt.Printf("input type is : %T", input)

	fmt.Println(reflect.TypeOf(input))
}
