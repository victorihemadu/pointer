package main

import "fmt"

type person struct {
	firstName string
	lastName string
	contactInfo
}

type contactInfo struct {
	email string
	postalcode int
}

func main() {
	jim := person{
		firstName: "Victor",
		lastName: "Ihemadu",
		contactInfo: contactInfo{
			email: "victorihemadu@gmail.com",
			postalcode: 100003,
		},
	}

	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = (newFirstName)
}

func (p person) print(){
	fmt.Printf("%+v", p)
}

