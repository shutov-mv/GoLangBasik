package main

import "fmt"

type contact struct {
	email string
	phone string
}

type person struct {
	name        string
	age         int
	contactInfo contact
}

func main() {
	var tom = person{
		name: "Tom",
		age:  24,
		contactInfo: contact{
			email: "tom@damain.local",
			phone: "88887775454",
		},
	}
	tom.contactInfo.email = "stom@damain.local"
	fmt.Println(tom.contactInfo.email)
	fmt.Println(tom.contactInfo.phone)
}
