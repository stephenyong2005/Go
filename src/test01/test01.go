package test01

import "fmt"

type Person struct {
	Name string
	Age int
	Address string
}

func PersonPrint(person Person){
	fmt.Println(person)
}
