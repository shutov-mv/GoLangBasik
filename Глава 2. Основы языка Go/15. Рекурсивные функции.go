package main

import "fmt"

func factorial(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
func main() {
	fmt.Println(factorial(4))
	fmt.Println(factorial(5))
	fmt.Println(factorial(6))
	fmt.Println(fibbonachi(4))
	fmt.Println(fibbonachi(5))
	fmt.Println(fibbonachi(6))
}

/*
Функция вычисляющая числа Фибоначи.
n-й ччлен последовательностей Фиьоначи
орпеделяется по формуле
f(n)=f(n-1) + f(n-2), причем f(0)=0, а f(1)=1.
*/
func fibbonachi(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibbonachi(n-1) + fibbonachi(n-2)
}
