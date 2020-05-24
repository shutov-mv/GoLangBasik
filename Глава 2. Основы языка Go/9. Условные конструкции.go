package main

import "fmt"

func main() {
	var (
		a = 5
		b = 6
	)
	if a > b {
		fmt.Println("верно")
	} else {
		fmt.Println("Не верно")
	}
	var (
		c float64 = 17
		d float64 = 17
	)
	if d > c {
		fmt.Println("Больше")
	} else if d < c {
		fmt.Println("Меньше")
	} else {
		fmt.Println("Ровно")
	}
}
