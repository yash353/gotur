package main

import (
	"bufio"
	"fmt"
	"os"

	"reflect"
)

func main() {

	fmt.Println("Enter the no..")

	a := bufio.NewReader(os.Stdin)

	input, _ := a.ReadString('\n')

	// input,error := a.ReadString('\n')

	// if error!= nil{

	// 	log.Fatal("Error in reading input")  	//to terminate program if error occur
	// }
	fmt.Println(reflect.TypeOf(input))
	// trimmedInput := str(input)
	// n,err := Strconv.parseInt(trimmedInput ,10, 32)

	// if err!= nil{
	// 	fmt.Println(n)
	// }

}
