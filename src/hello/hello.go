package main

import (
	"fmt"
	"test01"
	"strconv"
)




func main() {


	var person test01.Person
	person.Address = "Orlando, FL USA"
	person.Age = 25
	person.Name = "Stephen"

	test01.PersonPrint(person)

	fmt.Println(person.Name+" at "+strconv.Itoa(person.Age)+" years old, and he comes from "+person.Address)




}

