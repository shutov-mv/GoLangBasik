package main

import "fmt"

func main() {
	defer finish() //Оператор defer определяет функцию в конец программы
	fmt.Println("Program has been started")
	fmt.Println("Program is working")
	fmt.Println(divide(15, 5))
	fmt.Println(divide(4, 0))
	fmt.Println("stoped")
}
func finish() {
	fmt.Println("Program has been finished")
}
func divide(x, y float64) float64 {
	if y == 0 {
		panic("devision by zero")
	}
	return x / y
}
