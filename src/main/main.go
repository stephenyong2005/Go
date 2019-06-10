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

	var p = new(test01.Person)

	fmt.Println(p)

	//Testing for url fetching in the urlDemo package
	//urlDemo.FetchURL("http://www.bing.com")


	// Testing for web server in the webDemo package
	/*http.HandleFunc("/",webDemo.Handler)
	log.Fatal(http.ListenAndServe("localhost:8080",nil))
	*/
}



