package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// Address type
type Address struct {
	StreetAddress, City, Country string
}

// Person type
type Person struct {
	Name    string
	Friends []string
	Address *Address
}

// DeepCopy util copies the person object using Serialization
func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}     // Start with a new buffer
	e := gob.NewEncoder(&b) // Create an encoder with the buffer
	err := e.Encode(p)      // encode the person
	if err != nil {
		log.Panic(err)
	}
	d := gob.NewDecoder(&b) // Create a decoder on the bugger
	fmt.Println(string(b.Bytes()))
	result := Person{}      // prep the data type that will decode to
	err = d.Decode(&result) // decode into the preped data type
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(result)
	return &result
}

func main() {
	john := Person{"John",
		[]string{"Chris", "Matt", "Sam"},
		&Address{"123 London Rd", "London", "UK"}}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Jill")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
