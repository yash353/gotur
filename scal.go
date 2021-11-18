package main

import (
	"fmt"
	"sync"
	"time"
)

func square(a chan int) {

	defer wg.Done()
	// fmt.Println(<-a)
	c := <-a
	c = c * c
	fmt.Println("square is :", c)
	a <- c

}

func cube(a chan int) {
	c := <-a
	c = c * c * c
	a <- c
	fmt.Println("cubes is :", c)
	defer wg.Done()
}

func add(a chan int, b chan int, c chan int) {
	defer wg.Done()

	d := <-a
	e := <-b
	z := d + e
	c <- z
	time.Sleep(2 * time.Microsecond)
	// return

}

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	c := make(chan int)
	a := make(chan int)
	time.Sleep(2 * time.Microsecond)
	b := make(chan int)
	go square(a)
	// go cube(a)

	a <- 2
	go cube(b)
	b <- 2
	// b <- <-a
	// c <- 10
	// b <- <-a
	// fmt.Println(<-b)
	time.Sleep(2 * time.Microsecond)
	go add(a, b, c)
	fmt.Println("sum of squre and cubes :", <-c)
	wg.Wait()
}
