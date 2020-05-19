package main

import "fmt"

// Employee type
type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// NewEmployeeFactory << FUNCTIONAL APPROACH >> returns a factory function
func NewEmployeeFactory(postion string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, postion, annualIncome}
	}
}

func main() {
	// << FUNCTIONAL APPROACH >>
	developerFactory := NewEmployeeFactory("Developer", 150000)
	managerFactory := NewEmployeeFactory("Manager", 170000)

	john := developerFactory("John")
	jane := managerFactory("Jane")

	fmt.Println(john)
	fmt.Println(jane)
}
