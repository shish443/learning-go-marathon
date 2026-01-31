package main

import "fmt"

func main() {
	fmt.Println("__________________________")

	testMap := map[int]string{
		2:  "false",
		3:  "false",
		12: "true",
		5:  "true",
		1:  "true",
	}
	fmt.Println("testMap:", testMap) /*map умный он сортирует все или по алфавиту или по возрастанию*/
	testMap[3] = "false"             /*меняем значение мапа*/
	testMap[4] = "true"              /*добавляем значение в мап*/
	delete(testMap, 1)               /*удаляем значение из мапа*/
	delete(testMap, 10)              /*удаляем значение которого нет в мапе, ошибки не будет*/
	fmt.Println("testMap:", testMap)

	fmt.Println("__________________________")

	for something, bool := range testMap {
		fmt.Println("бесконечность не предел", something, bool) /*кстати в таком случае мап не сортирует вывод, а дает элементы по порядку*/
	}

	fmt.Println("__________________________")

	_, one := testMap[2]
	switch one {
	case true:
		fmt.Println("ваше значение верное")
	case false:
		fmt.Println("ваше значение НЕверное")
	}

	TestMap(testMap, 12, "iAmTired")
	fmt.Println("__________________________")
	fmt.Println("мап после функции", testMap)
	testNil()
}

func TestMap(revensMap map[int]string, revensInt int, revensBool string) { /*мап как и слайсы и тд передаются путями, а не дублируется, что прикольно*/
	revensMap[35] = "345srtg45w"
}

func testNil() {

	fmt.Println("-----------------------------------")

	/*тут будет nil-интересная штука, нулевое значение для мапов, слайсов и тд*/
	var testVarMap map[int]int

	fmt.Println("тут будет пусто", testVarMap)

	if testVarMap == nil {
		fmt.Println("я значения нил")
	} else {
		fmt.Println("я никогда не должен выполниться")
	}
	fmt.Println("-----------------------------------")

	/*почему я не заработаю?*/
	// testVarMap[1] = 1

	// fmt.Println("тут уже не будет пусто", testVarMap)

	// if testVarMap == nil {
	// 	fmt.Println("я значения нил")
	// } else {
	// 	fmt.Println("теперь я должен работать")
	// }

	/*вернее код заработает но вызовет панику, потому что мы не использовали маке для инициализации, как надо:*/

	testVarMap2 := make(map[int]string)

	if testVarMap2 == nil {
		fmt.Println("я значения нил")
	} else {
		fmt.Println("теперь я должен работать")
	}

	testVarMap2[2] = "something"
	fmt.Println("тут уже не будет пусто", testVarMap2)

	fmt.Println("-----------------------------------")
}
