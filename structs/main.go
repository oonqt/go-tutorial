package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo // same as contactInfo: contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// This type description means we work with a pointer , below this is the operator
// If we just pass in a person it will automatically make it a pointer to person for us
func (pointerToPerson *person) updateName(name string) {
	// (*pointerToPerson).firstName = name // Memory address that person exists at, it turn the pointer into a value
	(*pointerToPerson).firstName = name
}

func main() {
	alexa := person {
		firstName: "Alexa",
		lastName: "Marquez",
		contactInfo: contactInfo {
			email: "amarquez@memester.xyz",
			zipCode: 92562,
		},
	}

	// alexaPointer := &alexa // Gives the address this vediabel is pointing to it is a reference to the adddress the struct points to
	// alexaPointer.updateName("Alex")
	alexa.updateName("Alex")
	alexa.print()
}
