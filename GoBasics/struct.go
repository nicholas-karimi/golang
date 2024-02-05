/* A collection of Key value pair. Used to represent structured data
struct keys can hold any type not only primitive types such as int, uint, float, string, bool
type car Struct{
	Make string


}
*/

package main

import "fmt"

type messageToSend struct {
	// phoneNumber int
	message   string
	sender    user
	recepient user
}

type carMan struct {
	make   string
	model  string
	wheels int
	year   int
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" && mToSend.recepient.number == 0 {
		return false
	}
	if mToSend.sender.number == 0 && mToSend.recepient.name == "" {
		return false
	}
	return true
}

func test(msg messageToSend) {
	fmt.Printf("Sending message: '%s' to : %v\n", msg.message, msg.sender)
	fmt.Println("=========================================")
}

// Embeded struct
type car struct {
	model string
	make  string
}

type truck struct {
	car
	bedSize int
}

// Struct methods
type authenticationInfo struct {
	username string
	password string
}

// receiver
func (authInfo authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf(
		"Authorization: Basic %s: %s\n",
		authInfo.username,
		authInfo.password,
	)
}

func testAuth(authI authenticationInfo) {
	fmt.Println(authI.getBasicAuth())
	fmt.Println("==================")
}
func main() {
	test(messageToSend{
		// phoneNumber: 254710568014,
		sender: user{
			name:   "Nicholas Karimi",
			number: 254710568014,
		},
		recepient: user{
			name:   "onymous",
			number: 254788919010,
		},
		message: "Thanks for signing up with Textio!",
	})
	// test(messageToSend{
	// 	// phoneNumber: 254766524625,
	// 	sender: 254711568014,
	// 	recepient: "Michael John",
	// 	message:     "Thanks for signing up with Textio!",
	// })

	// Test struct
	fmt.Println("Struct...")
	mycar := carMan{
		make:   "Audi",
		model:  "SQ5",
		wheels: 4,
		year:   2017,
	}
	fmt.Println("My car...", mycar)
	fmt.Println("My car model us: ", mycar.model) // access struct field using the . operator

	fmt.Println("Starting Embeded struct....")
	lanesTruck := truck{
		bedSize: 10,
		car: car{
			model: "Subaru",
			make:  "XV",
		},
	}

	fmt.Println("Truck bed size: ", lanesTruck.bedSize)
	fmt.Println("Car make: ", lanesTruck.make)
	fmt.Println("Car model: ", lanesTruck.model)

	fmt.Println("Testing struct METHODS..")
	testAuth(
		authenticationInfo{
			username: "nkarimi",
			password: "41444141",
		},
	)
	testAuth(authenticationInfo{
		username: "jenkins",
		password: "2676hqfqy25",
	})
}

// fields in a structt can be access using the . operator

// Nested structs in GO
