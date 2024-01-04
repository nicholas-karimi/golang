package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Arrays in go...")
	fmt.Println("Lets bigin")

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100 // add to an array
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// b := [5]int // returns error -> [5]int (type) is not an expression
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)

	// array lenght
	fmt.Println("len:", len(a))

	/* Arrays are one-dimension - however you can create multi-dimensional arrays */
	var twoD [2][3]int
	for i :=0; i < 2; i++{
		for j :=0; j < 3; j++{
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

