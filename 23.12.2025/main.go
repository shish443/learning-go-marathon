package main

import "fmt"

var test1, test2, test3, test4 = true, 12, 12.574, "test" //о переменных объявленных вне функций знают все

func main() {
	test := "тест"

	fmt.Println("привет", test1, test2, test3, test4, test)
	//niam();variables();printf() так можно, но не желательно
	//так желательно:
	Niam()
	Variables()
	Printf()

}
func Niam() {
	fmt.Println("тест другой функции")
}
func Variables() { //variable-функция
	//var test1,test2,test3,test4 = true,12,12.574,test так нельзя, потому что функции о внутрянке друг о друга не знают
	var (
		test  = "test" //так как прошлая переменная test в main, то Variables об этом не знает и дает создать test
		test5 = ""
		test6 = 1
	)
	var test7 int
	//test8 string := так нельзя
	var test8 string //так можно
	test8 = "i am programmer box"

	fmt.Println(test, test1, test5, test6, test7, test8)
	fmt.Printf("%T\n", test5)
}
func Printf() {
	//test1,test2,test3,test4 = true,12,12.574,"test"

	fmt.Printf("%s что-то=%d,а если с дробной частью=%2f, ну и это конечно %t", test4, test2, test3, test1)
	//глаголы, так называются штуки после %
	//%v -обычное значение
	//%T-тип переменной
	//%d-целые цисла int
	//%f- с плавающей точкой float
	//%2f-отображает 2 символа после .
	//%s-для обучных строк
	//%q-строка в""
	//%t-для boolian (true,false) значений
}

//еще задался воспросом: как спизд... кхм взять переменную у другой функции, так вот это есть далее по курсу, так что скоро узнаю
