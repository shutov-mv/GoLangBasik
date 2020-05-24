package main

import "fmt"

func main() {
	var x int = 4
	var p *int
	p = &x
	fmt.Println(p)
	fmt.Println("Addres:", p)
	fmt.Println("Value:", *p)
	d := new(int)
	fmt.Println("vslue:", *d)
	*d = 8
	fmt.Println("value:", *d)
}
