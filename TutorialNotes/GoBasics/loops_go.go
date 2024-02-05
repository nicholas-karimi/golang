package main

import (
	"fmt"
)

/* given the cost threshold, maxMessages should calculate the max no of messages that can be sent */

func maxMessages(threshold float64) int {
	totalCost := 0.0

	// skip the condition because we dont know to what number we're looping upto
	for i := 0; ; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
		if totalCost > threshold {
			return i
		}
	}

}

func test(threshold float64) {
	fmt.Printf("Threshold: %.2f\n", threshold)
	max := maxMessages(threshold)
	fmt.Printf("Maximum messages that can be sent: = %v\n", max)
	fmt.Println("================================================================")
}

// easter egg. reward users if they send prime test messages this year.

func printPrimeMessages(max int) {
	for n := 2; n < max+1; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}
		if n%2 == 0 {
			continue
		}
		isPrime := true
		for i := 3; i*i < n+1; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if !isPrime {
			continue
		}
		fmt.Println(n)
	}
}

func testEE(max int) {
	fmt.Printf("Primes up to: %v:\n", max)
	printPrimeMessages(max)
	fmt.Println("================================================================")
}

func main() {
	test(10.00)
	test(20.00)
	test(100.00)
	fmt.Println("======Easter Egg Offer =====")
	testEE(10)
	testEE(11)
	testEE(102)

}
