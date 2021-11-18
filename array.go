package main

import "fmt"

func main() {

	// SINGLE DIMENSION ARRAY

	// var b [6] int = [6]int {5,3,5,3,4,6}

	// b:= []int {5,3,5,3,4,6}

	//b:= [...]int {5,3,5,3,4,6,7,4,5,7}

	// MULTI - DIMENSION ARRAY

	// var b= [3][3]int {{5,3,5},{3,4,6},{1,2,3}}

	// b := [3][3]int {{5,3,5},{3,4,6},{1,2,3}}

	b := [][]int{{5, 3, 5}, {3, 4, 6}, {1, 2, 3}}

	fmt.Println(b)
}
