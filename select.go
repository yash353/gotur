package main

import "fmt"

func sel(c, q chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-q:
			fmt.Println("q")
			return
		}
	}
}

func main() {
	c := make(chan int)
	q := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		q <- 0
	}()
	sel(c, q)
}
