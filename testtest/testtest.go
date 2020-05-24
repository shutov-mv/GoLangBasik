package main

import "fmt"

func main() {
	var (
		nambers  [5]int = [5]int{1, 2, 3, 4, 5}
		nambers1 [5]int = [5]int{1, 3}
		nambers2        = [...]int{1, 3}
		nambers3        = [...]int{1, 5, 7, 4, 3}
	)
	fmt.Println(nambers, nambers1)
	//var nambers1 [5]int = [5]int{1, 3}
	fmt.Println(nambers2, nambers3)
	nambers[0] = 54
	fmt.Println(nambers)
	fmt.Println(nambers[4])
	colors := [3]string{2: "Синий", 0: "Красный", 1: "Белый"}
	fmt.Println(colors[2])
	fmt.Println(colors)
}

]