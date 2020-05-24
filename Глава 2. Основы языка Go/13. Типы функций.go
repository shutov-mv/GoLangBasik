package main

import "fmt"

func add(x int, y int) int      { return x + y }
func multiply(x int, y int) int { return x * y }

func display(message string) {
	fmt.Println(message)
}

func main() {
	f := add
	fmt.Println(f(3, 4))

	f = multiply
	fmt.Println(f(3, 4))
	var f1 func(string) = display
	f1("Hello")
	action(10, 24, add)
	action(5, 6, multiply)
	slice := []int{-2, 4, 3, -1, 7, -4, 23}
	sumOfEvents := sum(slice, isEven)
	fmt.Println(sumOfEvents)
	sumOfPositive := sum(slice, isPositive)
	fmt.Println(sumOfPositive)
	ff := SelectFn(1)
	fmt.Println(ff(3, 4))
	ff = SelectFn(3)

	fmt.Println(ff(3, 4))
}

// 							Функция как параметр других функций
//
//
func action(n1 int, n2 int, operation func(int, int) int) {
	result := operation(n1, n2)
	fmt.Println(result)
}
func isEven(n int) bool {
	return n%2 == 0
}
func isPositive(n int) bool {
	return n > 0
}
func sum(numbers []int, criteria func(int) bool) int {
	result := 0
	for _, val := range numbers {
		if criteria(val) {
			result += val
		}
	}
	return result
}

func subtract(x int, y int) int {
	return x - y
}
func SelectFn(n int) func(int, int) int {
	if n == 1 {
		return add
	} else if n == 2 {
		return subtract
	} else {
		return multiply
	}

}
