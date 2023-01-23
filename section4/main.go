package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 94000,
		},
	}

	// go allows the underlying type to be passed to a receiver function that accepts a pointer
	alex.updateName("alexander")
	alex.print()
}

// struct is a value type. use a pointer to update it's values in another function
func (p *person) updateName(newFirstName string) {
	// go allows p.firstName instead of pointer de-reference (*p).firstName
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
