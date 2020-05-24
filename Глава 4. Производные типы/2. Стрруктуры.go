package main

import (
	"fmt"
	"strconv"
)

type person struct {
	name string
	age  int
}
type data struct {
	day       int
	month     int
	year      int
	FirstName string
	LastName  string
}

func main() {
	var tom = person{name: "Tom", age: 24}
	fmt.Println(tom.name)
	fmt.Println(tom.age)

	tom.age = 34
	fmt.Println(tom.name, tom.age)
	var test = data{day: 28, month: 7, year: 1989}
	var monthstr string = strconv.Itoa(test.month)
	if len(monthstr) == 1 { //len - возвраь длины стоки, масива
		monthstr = "0" + monthstr
	}
	var FullNeme = data{FirstName: "Ivan", LastName: "ivanov"}
	var датарождения = strconv.Itoa(test.day) + "." + monthstr + "." + strconv.Itoa(test.year)
	fmt.Println(датарождения)
	fmt.Println(FullNeme)
}
