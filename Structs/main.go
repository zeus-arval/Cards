package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	var artur person
	artur.firstName = "Vadim"
	fmt.Println(artur)
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Printf("%+v\n",alex)
	var eduard person
	fmt.Printf("%+v",eduard)
}
