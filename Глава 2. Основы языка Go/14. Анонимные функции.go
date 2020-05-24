package main

import "fmt"

func action(n1 int, n2 int, operatin func(int, int) int) {
	result := operatin(n1, n2)
	fmt.Println(result)
}
func selectFn(n int) func(int, int) int {
	if n == 1 {
		return func(x int, y int) int { return x + y }
	} else if n == 2 {
		return func(x int, y int) int { return x - y }
	} else {
		return func(x int, y int) int { return x * y }

	}
}
func square() func() int {
	var x int = 2
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := func(x, y int) int { return x + y }
	fmt.Println(f(3, 4))
	fmt.Println(f(6, 17))
	action(10, 25, func(x int, y int) int { return x + y })
	action(5, 6, func(x int, y int) int { return x * y })
	ff := selectFn(1)
	fmt.Println(ff(2, 3))
	fmt.Println(ff(4, 5))
	fmt.Println(ff(7, 6))
	fff := square()
	fmt.Println(fff())
	fmt.Println(fff())
	fmt.Println(fff())
}
