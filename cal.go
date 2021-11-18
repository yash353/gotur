package main

import (
	"fmt"
	"time"
)

func sq(a chan float64) {

	b := <-a

	fmt.Println("square of no.. is :", b*b)
	// time.Sleep(250 * time.Millisecond)
}

func cub(c chan float64) {

	s := <-c

	fmt.Println("cubes of no.. is :", (s*s)*s)
	// time.Sleep(250 * time.Millisecond)
}

func main() {

	a := make(chan float64)

	go sq(a)
	a <- 2.0
	time.Sleep(1 * time.Second)

	c := make(chan float64)

	go cub(c)
	c <- 2.0
	time.Sleep(1 * time.Second)

}
