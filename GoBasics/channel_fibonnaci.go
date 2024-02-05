/*  a sequence in which each number is the sum of two preceeding ones. 0, 1, 1, 2, 3, 5, 8, 13, 21, 34 */
package main

import (
	"fmt"
	"time"
)

func concurrentFibonnaci(n int) {
	chanInts := make(chan int)
	go func() {
		fibonnaci(n, chanInts)
	}()

	for v := range chanInts {
		fmt.Println(v)
	} 
}

// TEST SUITE
func test(n int) {
	fmt.Printf("Printing %v numbers...\n", n)
	concurrentFibonnaci(n)
	fmt.Println("======================")
}

func main() {
	test(10)
	test(5)
	test(20)
	test(13)

}

func fibonnaci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
}
