package main

import "fmt"

func main() {
	hello()
	hello()
	hello()
	add(4, 5)
	add(20, 6)
}
func hello() {
	fmt.Println("Hello World")
}
func add(x int, y int) {
	var z = x + y
	fmt.Println(z)

}
