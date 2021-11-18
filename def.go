package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(300 * time.Millisecond)

	for {

		select {

		case <-tick:
			fmt.Println(" --Yash--")

		case <-boom:
			fmt.Println(" -BOOM!!-")
			return

		default:
			fmt.Println("..........")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
