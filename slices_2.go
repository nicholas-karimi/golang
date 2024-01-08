package main

import "fmt"

// use variadic functions

func sum(nums ...int) int {
	num := 0
	for i := 0; i < len(nums); i++ {
		num += nums[i]
	}
	return num
}

func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func test(nums ...int) {
	total := sum(nums...)
	fmt.Println("Begin Variadic func================================================================")
	fmt.Printf("Summing %v costs...\n", len(nums))
	fmt.Printf("Bill for the month: %v\n", total)

}
func main() {
	fmt.Println("Spread Operator================================================================")
	names := []string{"bob", "nic", "joy"}
	printStrings(names...)
	fmt.Println("================================================================")
	// test(1, 2)
	fmt.Println("====== RANGE IN GO================================================================")
	fruits := []string{"apple", "orange", "banana", "mango"}
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}
}
