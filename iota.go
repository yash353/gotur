package main

import "fmt"

const (
	b = iota
	c = iota
	d = iota
)

const (
	x = iota
	_ //skip one value
	y
	_
	z
)

func main() {

	const a = iota

	fmt.Println(a)

	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}
