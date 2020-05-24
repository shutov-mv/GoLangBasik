package main

import (
	"fmt"
	"strconv" //Преобразование чисел в строки
)

func main() {
	var Имя = Person("Shutov", "Maxim")
	var date = Date(28, 07)
	fmt.Println(date)
	fmt.Println(Имя)
}
func Person(FerstName, LastName string) string {
	var ПолноеИмя = FerstName + " " + LastName
	return ПолноеИмя
}

func Date(Day, Mounth int) string {
	var ДатаРождкние = strconv.Itoa(Day) + "." + strconv.Itoa(Mounth) //strvconv.Ita - конвертация типов данных
	return ДатаРождкние
}
