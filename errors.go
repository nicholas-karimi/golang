

package main 

import "fmt"


func sendSMSToCouple(msgToCustomer, msgToSpouse string)(float64, error){
	costForCustomer, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, nil
	}
	costForSpouse, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0.0, nil
	}

	return costForCustomer + costForSpouse, nil
}

func sendSMS(message string) (float64, error){
	const maxLength = 25
	const costPerChar = 0.0002

	if len(message) > maxLength{
		return 0.0, fmt.Errorf("can't send texts over %v characters", maxLength)
	}
	return costPerChar * float64(len(message)), nil
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
func main(){
	test("I will place an order next week. Please reserve the Item.", "This looks amazing!")
	test("Thank you for shopping with us.", "TEnjoy rest of your evening!")
	test("Lend me your wife.", "No!")

}