package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := pesron{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
}
