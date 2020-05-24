package main

import "fmt"

func main() {
	var people = map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 8,
	}
	fmt.Println(people)
	fmt.Println(people["Alice"])
	fmt.Println(people["Bob"])
	people["Bob"] = 32
	fmt.Println(people["Bob"])
	if val, ok := people["Tom"]; ok {
		fmt.Println(val)
	}
	for kye, value := range people {
		fmt.Println(kye, value)
	}
}
