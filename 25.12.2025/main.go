package main

//константы

import "fmt"

func main() {
	// несколько констант или/и функций можно объявлять через ()
	// константам сразу нуножно присваивать значения
	// в отличии от var const можно объявить и не использовать
	const (
		test1 = 1                     // int
		test2 = "i am programmer box" //string
		test3 = 5.7372627             //float64
		test4 = true /*bool*/)
	// вот так: /*коментарий*/ теперь всегда делать
	// ---------
	// начинает с 0
	// const (a= iota
	// b
	// c
	// d)
	// -----	-----------
	// но может начинать с 1:
	const (
		a = iota + 1
		b
		c
	)

	fmt.Printf("понедельник по счету %d день недели \n", a)
	//----------------
	var (
		test5 int     = 926
		test6 string  = "i am test various"
		test7 float64 = 1535.000142
		test8 bool    = false
		test9 int     = 4
	)
	//----------------
	fmt.Println(test7)
	test7 = test1 + test7 //полная запись
	fmt.Println(test7)
	test7 += test3 //краткая запись
	fmt.Println(test7)
	//а так я могу?
	// нет test7 =- test1, но могу так:
	test7 -= test1
	fmt.Println(test7)
	test5 *= test9
	fmt.Println(test5)
	fmt.Println(test5, test6, test7, test8, test9)
}
