/* Go has no explicit implementation of while loop with the while keyword. Instead it uses the for keyword that omits the inital and agfter statements */

package main

import "fmt"

func main() {
	age := 17

	for age < 18 {
		fmt.Printf("still a teenager at %d\n", age)
		age++


	}
	fmt.Println("you'll be an adult in %v years", age)
}
