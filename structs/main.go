package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}

func main() {
	// alex := person{firstname: "Alex", lastname: "Anderson"}
	// fmt.Println(alex)
	// var alex person
	// alex.firstname = "Alex"
	// alex.lastname = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{firstname: "Jim", lastname: "Halls", contact: contactInfo{email: "jim@gmail.com", zipCode: 95347}}
	fmt.Printf("%+v", jim)
}
