package main

import "fmt"

func main() {
	testIf()
	testIfElse()
	multipleConditions()
	statementsInIf()
	testSwitch()
}

func testIf() {
	testInt := 10
	testFloat64 := 5.00
	if testInt > int(testFloat64) /*значение должно быть типа булиан, то есть true/false */ {
		fmt.Println("вам бонус")
	}
	if testFloat64 > 2 {
		fmt.Println("test send")
	}

	testFloat64 = 1.00

	if testFloat64 > 2 {
		fmt.Println("test send")
	}
	fmt.Println("_________________________")
}

func testIfElse() {
	testString := "i love Korora"
	if testString == "i love Korora" {
		fmt.Println("you are Nikita")
	} else {
		fmt.Println("you aren't Nikita")
	}

	fmt.Println("_________________________")

	timeForDrinkCoffe := 10
	timeForDrinkTea := 3
	if timeForDrinkTea > 30 || timeForDrinkCoffe > 30 {
		fmt.Println("you need go to lanch")
	} else if timeForDrinkCoffe <= 5 || timeForDrinkTea <= 5 {
		fmt.Println("you are speed")
	} else if timeForDrinkCoffe < 30 || timeForDrinkTea < 30 {
		fmt.Println("you need do your work after drink break")
	} else /*этой строчки модет не быть. код просто ничего не выполнит*/ {
		fmt.Println("you need go office")
	}
	/*&&-и ||-или*/
	fmt.Println("_________________________")
}

func multipleConditions() {
	var (
		testInt1    = 1
		testInt2    = 2
		testString  = "i am ok"
		testBool    = false
		testfloat64 = 0.00
	)
	/*reccomend cod if you don't need do something else*/
	if (testInt1 == 1) && (testInt2 == 2) && testBool && testfloat64 >= 1.0 /*()-не обязательны и без них работает, но они повышают читаемость кода*/ {
		fmt.Println("you have a normal programmer")
	} else {
		fmt.Println("you don't have a normal programmer,because you need hire me for a job")
	}

	fmt.Println("_________________________")

	/*not recommend if you don't need do something else*/
	if testString == "i am ok" {
		if testInt1 == 1 {
			if testInt2 == 2 {
				if testfloat64 <= 0.0 {
					fmt.Println("this expression is correct", testBool)
				} //это есле я должен продублировать на каждое иф
			} //это есле я должен продублировать на каждое иф
		} //это есле я должен продублировать на каждое иф
	} else {
		fmt.Println("you don't hsve a normal programmer")
	}

	fmt.Println("_________________________")
}

func statementsInIf() {
	fucnForTest := func(test int) int {
		test = test * test
		return test
	}

	testYourInt := 2345
	//синтаксическт правильно, но логически всегда не правильно:
	// if tegaudsfigsayerfuv, testVarForIf := fucnForTest(testYourInt), 183; tegaudsfigsayerfuv >= 10 && testVarForIf <= 8 {
	// 	fmt.Println("you have a test messenge", tegaudsfigsayerfuv, testVarForIf)
	// }

	if tegaudsfigsayerfuv, testVarForIf := fucnForTest(testYourInt), 183; tegaudsfigsayerfuv >= 10 && testVarForIf <= 1000 {
		fmt.Println("you have a test messenge", tegaudsfigsayerfuv, testVarForIf)
	}

	fmt.Println("_________________________")
}
func testSwitch() {
	const command string = "test"
	switch command {
	case "start":
		fmt.Println("стрим начился")
	case "stop":
		fmt.Println("стрим закончился")
	case "info":
		fmt.Println("информация о стриме: стрим по минекрафту")
	case "ban":
		fmt.Println("пользователь забанен")
	case "game":
		fmt.Println("введите название предложенной игры")
	default:
		fmt.Println("неизвестная команда")
	}
}
