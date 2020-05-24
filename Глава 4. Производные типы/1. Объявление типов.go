package main

import "fmt"

type mile int
type kilometr int
type library []string

func printBooks(lib library) {
	for _, value := range lib {
		fmt.Println(value)
	}
}

func distanceToEnemy(distance mile) {
	fmt.Println("Расстояние до противника:")
	fmt.Println(distance, "миль")
}

func main() {
	var distance mile = 5
	distanceToEnemy(distance)
	var distance2 kilometr = 5
	fmt.Println(distance)
	distance += 5
	fmt.Println(distance)
	fmt.Println(distance2)
	var myLibrary library = library{"Books1", "Books2", "Books3", "Books4"}
	printBooks(myLibrary)

}
