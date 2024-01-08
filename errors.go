package main

import (
	"fmt"
	"strconv"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	costForCustomer, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}
	costForSpouse, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0.0, err
	}

	return costForCustomer + costForSpouse, nil
}

func sendSMS(message string) (float64, error) {
	const maxLength = 25
	const costPerChar = 0.0002

	if len(message) > maxLength {
		return 0.0, fmt.Errorf("can't send texts over %v characters", maxLength)
	}
	return costPerChar * float64(len(message)), nil
}

// formating strings
func getSMSErrorString(coast float64, recepient string) string {
	return fmt.Sprintf("SMS that costs $%.2f  cannot be sent to %s", coast, recepient)
}

func test(msgToCustomer, msgToSpouse string) {
	defer fmt.Println("=============================")
	fmt.Println("Message to customer: ", msgToCustomer)
	fmt.Println("Message to spouse: ", msgToSpouse)

	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Total cost $", totalCost)

}
func testSMSError(cost float64, recepient string) {
	s := getSMSErrorString(cost, recepient)
	fmt.Println(s)
	fmt.Println("=============================")
}
func main() {
	test("I will place an order next week. Please reserve the Item.", "This looks amazing!")
	test("Thank you for shopping with us.", "TEnjoy rest of your evening!")
	test("Lend me your wife.", "No!")


	fmt.Println("Formated error string msg............")
	testSMSError(1.4, "Kenny")
	testSMSError(2.1, "+254 710 568 014")
	testSMSError(0.01, "Johnson")

	fmt.Println("Test ASCII to String conversion with err handling.....")
	i, err := strconv.Atoi("42b")
	if err != nil {
		fmt.Println("Couldn't convert:", err) // Couldn't convert: strconv.Atoi: parsing "42b": invalid syntax
		return
	} else {
		fmt.Println("Converted string is :", i)
	}

}
