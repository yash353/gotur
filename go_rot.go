package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("\t Hello \n its run go--routine")
}
func main() {
	go hello()
	go numbers()
	go alphabets()
	time.Sleep(1 * time.Second)
	fmt.Println("\n main function terminated")
	// fmt.Println(runtime.GOMAXPROCS(0))
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
