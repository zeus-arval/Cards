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
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Anderson",
		age: 15,
		contact: contactInfo{
			email: "jim.anderson@gmail.com",
			zip: 94000,
		},
	}
	fmt.Printf("%+v", jim)
}
