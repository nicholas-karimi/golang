package main

import "fmt"

func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

type email struct {
	message     string
	fromAddress string
	toAddress   string
}

func (e email) print() {
	fmt.Println("message:", e.message)
	fmt.Println("fromAddress:", e.fromAddress)
	fmt.Println("toAddress:", e.toAddress)

}

func test(e *email, newMessage string) {
	fmt.Println("--before --")
	e.print()
	fmt.Println("-- end before--")
	e.setMessage("this is my second draft")
	fmt.Println("-- after--")
	e.print()
	fmt.Println("-- end after--")
	fmt.Println("=================================")
}
func main() {
	test(&email{
		message:     "this is my first draft",
		fromAddress: "kqO2N@example.com",
		toAddress:   "l5Yp4@example.com",
	}, "this is my second draft")
	test(&email{
		message:     "this is my fourth draft",
		fromAddress: "ken@example.com",
		toAddress:   "l5Yp4@example.com",
	}, "this is my really fourth draft")
	test(&email{
		message:     "this is my thired draft",
		fromAddress: "mike@example.com",
		toAddress:   "yp4@example.com",
	}, "this is my hahah draft")
}
