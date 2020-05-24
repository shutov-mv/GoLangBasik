package main

import "fmt"

func main() {
	var users = []string{"Tom", "Alice", "Kate"}
	fmt.Println(users)
	fmt.Println(users[2])
	for _, value := range users {
		fmt.Println(value)
	}
	var user []string = make([]string, 3)
	user[0] = "Tomara"
	user[1] = "Perova"
	user[2] = "Ivanov"
	fmt.Println(user)
	users = append(users, "Bob")
	for _, value := range users {
		fmt.Println(value)
	}
}
