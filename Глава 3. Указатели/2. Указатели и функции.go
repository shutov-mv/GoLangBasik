package main

import "fmt"

func changeValue(x int) {
	x = x * x
}
func changeValue1(x *int) {
	*x = (*x) * (*x)
}
func createPointer(x int) *int {
	p := new(int)
	*p = x
	return p
}

func main() {
	d := 5
	fmt.Println("d before:", d)
	changeValue(d)
	fmt.Println("d after:", d)
	d1 := 5
	fmt.Println("d before:", d1)
	changeValue1(&d1)
	fmt.Println("d after:", d1)
	p1 := createPointer(7)
	fmt.Println("p1:", *p1)
	p2 := createPointer(10)
	fmt.Println("p2:", *p2)
	p3 := createPointer(28)
	fmt.Println("p2:", *p3)
}
