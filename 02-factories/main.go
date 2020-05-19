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

// EmployeeFactory << STRUCTURAL APPROACH >>
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

// Create returns a pointer to a newly created employee
func (e *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, e.Position, e.AnnualIncome}
}

func main() {
	// << FUNCTIONAL APPROACH >>
	developerFactory := NewEmployeeFactory("Developer", 150000)
	managerFactory := NewEmployeeFactory("Manager", 170000)

	john := developerFactory("John")
	jane := managerFactory("Jane")

	fmt.Println(john)
	fmt.Println(jane)

	// << STRUCTURAL APPROACH >>
	ceoFactory := EmployeeFactory{"CEO", 200000}

	sam := ceoFactory.Create("Sam")
	fmt.Println(sam)
}
