package main

import "fmt"

type A struct {
	age int

	name string
}

func main() {
	a := A{

		// 55,

		// "yash",

		age:  55,
		name: "yash",
	}

	fmt.Println(a)
	fmt.Println(a.name)
}
