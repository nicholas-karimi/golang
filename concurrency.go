package main

import (
	"fmt"
	"time"
)

func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

func test(message string) {
	sendEmail(message)
}
func main() {
	test("Hi there Nick")
	test("Hello Kenny")
	test("Hello there Paul Kenny")
	test("Hi there Mikey")
	test("Bonjour Maureen")

}
