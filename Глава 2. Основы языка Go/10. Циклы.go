package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println(i * i * i)
	}
	for a := 1; a <= 10; a++ {
		for b := 1; b <= 10; b++ {
			fmt.Print(a*b, "\t")
		}
		fmt.Println()
	}

	var users = [3]string{"Вася", "Петя", "Таня"}
	for index, value := range users {
		fmt.Println(index, value)
	}
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}
}
