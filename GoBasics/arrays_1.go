package main

import "fmt"

func getMessageWithReplies() [3]string {
	return [3]string{
		"click here to sign up",
		"pretty please click here",
		"we beg you to sign up",
	}
}

func send(name string, doneAt int) {
	fmt.Printf("sending to %v...", name)
	fmt.Println()

	messages := getMessageWithReplies()
	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf(`sending: "%v"`, msg)
		fmt.Println()

		if i == doneAt {
			fmt.Println("they responded!")
			break
		}
		if i == len(messages)-1 {
			fmt.Printf("complete failure")
		}
	}
}

func main() {
	send("Victor", 0)
	send("Hezreon", 10)
	send("Mike", 1)
	send("John", 3)
}
