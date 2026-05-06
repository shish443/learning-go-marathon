package main

import (
	"fmt"
) // Импорт внешней библиотеки на айпаде сделать не могу, приеду домой поэкспериментировать с внешними библиотеками
func main() {
	fmt.Println("hello", "kak dela")
	// автоматическое определение
	var test1 = "server is working"
	//кратко обьявить переменную
	test3 := "server is тест working"
	// int
	//string
	//bool правда ложь
	//float дробные числа, то есть с ,
	var test2 int = 1
	var test4 float64 = 5.73635373839027372529
	//Булево значение (Boolean) — это < логический тип данных, который может принимать только два состояния: истина (true) или ложь (false).
	var isGoEasy bool = true //Истина
	var isTired bool = false //Ложь
	fmt.Println(isGoEasy, isTired, test1, test2, test3, test4)

	// арифметика
	var a float64 = 3
	var b float64 = 65

	summ := a + b
	vichitanie := b - a
	umnohenie := b * a
	delenie := b / a
	var a2 = 3
	var b2 = 65
	delenie2 := b2 / a2
	var OstatokOtDelenia int = b2 % a2 //остаток от деления
	const мойденьроддения = "7 апреля" //константу нельзя изменять

	fmt.Println("сумма:", summ, ", вычитание:", vichitanie, ", умножение:", umnohenie, delenie, delenie2, OstatokOtDelenia, мойденьроддения)

}
