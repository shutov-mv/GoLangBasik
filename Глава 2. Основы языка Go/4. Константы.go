package main

import "fmt"

func main() {
	const pi float64 = 3.1415
	fmt.Println(pi)
	const (
		fp float64 = 7.56687
		dd float64 = 45.2395
	)
	fmt.Println(fp, dd)
	const (
		a = 3
		b = 2
		c = 9
		d = 2.3
		i = 3.5
	)
	fmt.Println(a, b, c, d, i)

}
