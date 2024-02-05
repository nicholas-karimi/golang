/* range iterates over elements in a variety of data structures.
range on strings iterates over Unicode code points. The first value is the starting byte index of the rune and the second the rune itself
*/

package main

import (
	"fmt"
)

func main() {

	s := []int{2, 3, 5, 7, 11, 13}
	var p = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(p)
	for idx, v := range p { // idx is index, v is value
		fmt.Printf("%v -> %v\n", idx, v)
	}


	// cal sum
	sum := 0
	for _, v := range s { // ignore index using the _ identifier
		sum += v
	}
	fmt.Println("Total sum for is: ", sum)

	// iterate over keys only
	for k := range p {
		fmt.Println("Key: ", k)
	}

	// range on strings ->  The first value is the starting byte index of the rune and the second the rune itself.
	for idx, char := range "mamamboga" {
		fmt.Println(idx, char)
	}

	// msg
	const name = "Nicholas Karimi"
    const openRate = 100.5

    msg := fmt.Sprintf("Hi %s, your rate is %.1f percent", name, openRate)
	fmt.Println(msg)
}
