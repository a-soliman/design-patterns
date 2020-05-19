package main

import "fmt"

// Address type
type Address struct {
	zip, street string
}

// User type
type User struct {
	name  string
	age   int
	phone string
	Address
}

func (u *User) String() string {
	return fmt.Sprintf("Name: %s, is %d, phone #: %s\nLives at %s, %s\n", u.name, u.age, u.phone, u.street, u.zip)
}

// UserBuilder type
type UserBuilder struct {
	user *User
}

// NewUserBuilder returns a pointer to a userBuilder
func NewUserBuilder() *UserBuilder {
	return &UserBuilder{&User{}}
}

// Build return a pointer to a user of a user builder
func (u *UserBuilder) Build() *User {
	return u.user
}

// SetName sets the name of a given userBuilder
func (u *UserBuilder) SetName(name string) *UserBuilder {
	u.user.name = name
	return u
}

// SetAge sets the age of a given userBuilder
func (u *UserBuilder) SetAge(age int) *UserBuilder {
	u.user.age = age
	return u
}

// SetPhone sets the phone of a given userBuilder
func (u *UserBuilder) SetPhone(phone string) *UserBuilder {
	u.user.phone = phone
	return u
}

// SetAddress sets the address of a given userBuilder
func (u *UserBuilder) SetAddress(zip, street string) *UserBuilder {
	u.user.Address = Address{zip, street}
	return u
}

func main() {
	user := NewUserBuilder().
		SetName("Zaki").
		SetAge(50).
		SetAddress("94999", "Main St.").
		Build()
	fmt.Print(user.String())
}
