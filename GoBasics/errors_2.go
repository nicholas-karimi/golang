package main

import (
	"errors"
	"fmt"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0.0, errors.New("Can't divide by zero")
	}
	return x / y, nil
}

func test(x, y float64) {
	defer fmt.Println("============================================")
	fmt.Printf("Dividing %.2f by %.2f ...\n", x, y)
	quotient, err := divide(x, y)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Quotient: %.2f", quotient)
	fmt.Printf("%.2f / %.2f = %.2f\n", x, y, quotient)

}

func main() {
	test(10, 0)
	test(10, 2)
	test(100, 7)
	test(97, 4)
	test(15, 30)
}
