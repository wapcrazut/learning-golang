package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	p := person{
		firstName: "John",
		lastName:  "Smith",
		contactInfo: contactInfo{
			email:   "test@test.com",
			zipCode: 9400,
		},
	}

	// Alternative pointer method
	//pointer := &p
	//pointer.updateName("Pointer")

	p.updateName("Zenon")
	p.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(name string) {
	(*p).firstName = name
}
