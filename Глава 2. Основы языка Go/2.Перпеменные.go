/*

	ПерваЯ программа


*/

package main //Определение пакета для текущего файла

import "fmt" //Подключение пакета fmt

func main() { //Определение функции
	//fmt.Println("Hello Go!")	//Выврд строки в консоль
	var hello string
	hello = "Hello World"
	fmt.Println(hello)
	/*
		var (
			name string = "Maxim"
			age  int    = 30
		)
		fmt.Println(name)
		fmt.Println(age)
		var hello string = "Hello world"
	*/
	fmt.Println(hello)
	hello = "Hello Go"
	fmt.Println(hello)
	hello = "Привэдъ Медвэдъ"
	fmt.Println(hello)
	name := "Tom"
	age := 30
	fmt.Println(name, age)
}
