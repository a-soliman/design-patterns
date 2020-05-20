package main

import "fmt"

// Address type
type Address struct {
	StreetAddress, City, Country string
}

// DeepCopy makes a deep copy of an address object
func (a *Address) DeepCopy() *Address {
	return &Address{
		a.StreetAddress,
		a.City,
		a.Country,
	}
}

// Person type
type Person struct {
	Name    string
	Address *Address
	friends []string
}

// DeepCopy makes a deepCopy of a person
func (p *Person) DeepCopy() *Person {
	q := *p
	q.Address = q.Address.DeepCopy()
	copy(q.friends, p.friends)
	return &q
}

func main() {
	john := Person{
		"John",
		&Address{"123 London Rd", "London", "UK"},
		[]string{"Jack", "Mike"},
	}
	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "132 Baker St"
	jane.friends = append(jane.friends, "Angela")
	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
