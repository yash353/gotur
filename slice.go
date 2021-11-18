package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} //slice

	b := a[:] //copy a into b

	c := b[1:] //print from index one

	d := b[:7] //print from 0 index to 6

	fmt.Println(a)

	fmt.Println(b)

	fmt.Println(c)

	fmt.Println(d)
}
