/* Interfaces are a collection of method signatures */

package main

import (
	"fmt"
)

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlypay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlypay * c.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

// test
func testData(e employee) {
	fmt.Println(e.getName(), e.getSalary())
	fmt.Println("=============================")
}

func main() {
	testData(fullTime{
		name:   "Kenny Paul",
		salary: 260000,
	})
	testData(contractor{
		name:         "Jill Paul",
		hourlypay:    160,
		hoursPerYear: 700,
	})
	testData(contractor{
		name:         "Nicholas Karimi",
		hourlypay:    100,
		hoursPerYear: 1000,
	})

}
