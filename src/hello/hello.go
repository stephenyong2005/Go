package main

import (
	"fmt"

	"stringutil"
	"strconv"
)

type Person struct {
	Name string
	Age int
	Address string
}

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))

	const (
		a = 10
		b
		c
		d
	)

	fmt.Println(a,b,c,d)



	var person Person
	person.Address = "Orlando, FL USA"
	person.Age = 25
	person.Name = "Stephen"

	fmt.Println(person.Name+" at "+strconv.Itoa(person.Age)+" years old, and he comes from "+person.Address)

	var s = 200

	p := &s
	fmt.Println(*p)

	person.AgePlus(5)

	fmt.Println(person.Name+" at "+strconv.Itoa(person.Age)+" years old, and he comes from "+person.Address)





}

func (p *Person) AgePlus(n int) int {


	p.Age = p.Age + n
	fmt.Println(p.Age)

	return 0

}