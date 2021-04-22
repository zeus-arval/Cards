package main

import "fmt"

type contactInfo struct {
	zip   int
	email string
}

type person struct {
	firstName string
	lastName  string
	age       int
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Anderson",
		age: 15,
		contactInfo: contactInfo{
			email: "jim.anderson@gmail.com",
			zip: 94000,
		},
	}
	jim.print()
}

func (p person) print(){
	fmt.Printf("%+v", p)
}