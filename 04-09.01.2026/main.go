package main

import (
	"fmt"
	"slices"
)

func main() {
	sleceForDeleteElementInSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	array()
	clice()
	changeSlice()
	onlySlice()
	testCapacity()
	testMakeSlice()
	testDeleteElementInSlice()
	deleteElementInSlice(2, 4, sleceForDeleteElementInSlice)
	testFor()
}

func array() {
	pointerTest := "testLine"
	// testPointer := &pointerTest
	fmt.Println(len(pointerTest))

	var test [10]int
	fmt.Println(test)

	test[4] = 124
	fmt.Println(test)

	test2 := [6]int{0, 1, 2, 3, 4, 5}
	fmt.Println(test2[3])
	/*len выводит общее количсектво элементов в массиве, строке и тд */
	fmt.Println(len(test2))

	test2[len(test2)-1] = 235216
	fmt.Println(test2[5])

	test2[len(test2)-3] = 555 /*обратиться к конкретному значению массивая через len я не могу, так как len выводит общее количсектво элементов. Можно сделать через кастыль*/
	fmt.Println(test2[3])
	fmt.Println("_______________________________")

}

func clice() {
	testArray := [6]int{0, 1, 2, 3, 4, 5}
	testSlice := testArray[:]
	testSlice = testArray[1:5]

	fmt.Println(testArray)
	fmt.Println(testSlice)

	testSlice = testArray[1:3]
	fmt.Println(testSlice)
	fmt.Println("_______________________________")

}

func changeSlice() {
	newArray := [11]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(newArray)

	newSlice := newArray[1:4]
	fmt.Println(newSlice)
	newSlice[2] = 111
	fmt.Println(newSlice)
	fmt.Println(newArray)
	fmt.Println("_______________________________")
}

func onlySlice() {
	slice := make([]int, 10)
	fmt.Println(slice)
	slice = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice)
	fmt.Println("_______________________________")
}

func testCapacity() {
	var testArray = [5]string{"1", "Cake", "3", "4", "5"}
	testArray[1] = "test1213"
	fmt.Printf("место в памяти:%p \n", &testArray)
	fmt.Println("массив testArray", testArray, "лен, то есть сколько сейчас в массиве testArray", len(testArray))
	fmt.Printf("======================\n\n")

	slicetestArray := testArray[1:5]
	AddressSliceTestArray := &slicetestArray

	textToPrintf := func() {
		fmt.Printf("место в памяти массива testArray:%p и место в памяти слайса slicetestArray:%p и место в памяти для\"cake\"в массиве:%p и в слайсе:%p\n", &testArray, &slicetestArray, &testArray[1], &slicetestArray[1])
		fmt.Println("массив testArray", testArray)

		fmt.Printf(
			"Слайс: %v\n"+
				"Len исходного testArray: %d\n"+
				"Cap слайса: %d | Len слайса: %d\n"+
				"Адрес самого СЛАЙСА (пульта): %p\n"+
				"Адрес СКРЫТОГО МАССИВА (куда смотрит пульт): %p\n"+
				"======================\n\n",
			slicetestArray,        // %v
			len(testArray),        // %d
			cap(slicetestArray),   // %d
			len(slicetestArray),   // %d
			AddressSliceTestArray, // %p (это адрес переменной-слайса)
			&slicetestArray[0],    // %p (это адрес первого элемента = адрес массива)
		)
	}
	fmt.Println("слайс:", slicetestArray, "лен массива testArray:", len(testArray), "cap, то есть максимальная длинна слайса slicetestArray:", cap(slicetestArray), "лен слайса slicetestArray:", len(slicetestArray))
	textToPrintf()

	slicetestArray = append(slicetestArray, "testAppendString")
	textToPrintf()
	/*лен (len)-сколько сейчас в массиве, слайсе. кап(cap)-сколько максим слайса*/

	slicetestArray = append(slicetestArray, "testReplaceSlice")
	textToPrintf()

	slicetestArray = append(slicetestArray, "secondTestAppendString")
	textToPrintf()

	slicetestArray = append(slicetestArray, "thirdTestAppendString")
	textToPrintf()

	slicetestArray = append(slicetestArray, "fourfTestAppendString")
	textToPrintf()

	slicetestArray = append(slicetestArray, "test1")
	textToPrintf()

	slicetestArray = append(slicetestArray, "test2")
	textToPrintf()

	fmt.Println("_______________________________")
}

func testMakeSlice() {
	slice := make([]int, 0, 1000)
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Printf("Массив:%v\nНынешнее количество:%d \nМаксимальное количество:%d\n", slice, len(slice), cap(slice))
	fmt.Println("_______________________________")
}
func testDeleteElementInSlice() {
	/*массивы практически не испольщуются и в основном используются слайсы*/
	/*если в [] есть значение-это массив, если пусто-это слайс*/
	test2DeleteElementInSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(test2DeleteElementInSlice)
	test2DeleteElementInSlice = append(test2DeleteElementInSlice[:2], test2DeleteElementInSlice[2+1:]...)
	fmt.Println(test2DeleteElementInSlice)
	fmt.Println("_______________________________")
}
func deleteElementInSlice[T any](index int, index2 int, slice []T) []T /*это означает,что я ввозвращаю любое, в данном случае слайс*/ {
	//[T any]-это дженерики (Generics), они нужны, что бы говорить, что это может быть любой тип(int,bool,strig,float etc)
	//пример чере append:testDeleteElementInSlice = append(elementInSlice[:2], elementInSlice[2+1:]...)
	fmt.Println(slice)
	slice = slices.Delete(slice, index, index+index2)
	fmt.Println(slice)
	fmt.Println("_______________________________")
	return slice
}
func testFor() {
	testslice := []string{"test1", "test2", "test3", "test4"}
	for _, name := range testslice {
		fmt.Printf("%s\n", name)
	}
}
