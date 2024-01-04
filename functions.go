/*
function signature
func sub(x int, y int)int {
	return x-y
}
*/

package main

import (
	"errors"
	"fmt"
)

func addNumbers(a int, b int) int {
	return a + b
}

// concat string
func concat(s1 string, s2 string) string {
	return s1 + s2
}

/*
go support syntatic sugar i.e when multiple arguements are of the same type, the type only needs to be declared on the last argument
func concat (s3, s4 string) string {}
*/
/*
Go func callbacks
f func(func(string, string)string, string) string
-> whenever we pass a calback in Go the type of  the function changes based on what its inputs and outputs are.
*/

// pass by value
func increment(x int) int {
	res := x + 1
	return res
}

// ignotr return values
func getNames() (firstName, lastName string) {
	// return "John", "Doe" -> explicit return- overiwrites the return value
	firstName = "John"
	lastName = "Doe"
	return
}

// Early Return
func divide(divinded, divisor int) (int, error) {
	divinded = 10
	divisor = 6
	if divisor == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return divinded / divisor, nil
}

func main() {
	res := addNumbers(1, 2)
	fmt.Println("Result: ", res)

	// concat str
	fmt.Println(concat("Hello ", "World!"))

	// increment
	x := 9
	x = increment(x)

	fmt.Println("Check pass by value mutation possibility. X is: ", x)

	// ignore return value and variables

	firstName, _ := getNames()
	fmt.Println("Welcome to Textion: ", firstName)

	fmt.Println("Early Returns")
	fmt.Println("Division by: ", divide(10, 0))
}
