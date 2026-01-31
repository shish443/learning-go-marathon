package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 3; i >= 1; i-- {
		fmt.Println("name int:", i)
	}
	fmt.Println("_________")

	//бесконецный цикл:
	// for{
	//fmt.Println("_________")
	// }

	seconds := 2
	//time-функция которая говорит приостановить на н секунд, минут, часов выполнение программы
	for seconds > 0 {
		println("wait", seconds, "seconds")
		time.Sleep(1 * time.Second)
		seconds--
	}
	fmt.Println("_________")
	тестРусский()
	breakAndContinue()
	ForRangeBreakContinue()
}

func тестРусский() {
	fmt.Println("Эта функция называеся на русском")
	fmt.Println("_________")
}

func breakAndContinue() {
Loop: /*название цикла*/
	for {
		var mut int
		fmt.Print("на сколько вы хотите пойти понюхать траву?(возможно от 1 до 10 минут, для выхода введите: 0)")
		fmt.Scan(&mut)

		switch mut {
		case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
			fmt.Println("вы трогаете траву:", mut, "минут")
			time.Sleep(time.Duration(mut) * time.Minute)
		case 0:
			fmt.Println("вы не трогаете траву, фуу задрот")
			break Loop /*мы указываем откуда именно выходим, так как по умолчанию он выйдет из switch*/
		default:
			fmt.Println("неверное значение")
			continue Loop /*мы указываем откуда именно выходим, так как по умолчанию он выйдет из switch*/
		}
	}
}

func ForRangeBreakContinue() {
	dailyIncome := [] /*7-тогда массив*/ int{136, 262, 893, 439, 5490, 22246, 677}
	daysOfWeek := []string{"m", "t", "w", "t", "f", "s", "s"}
	var totalIncome int

	for i, v := range dailyIncome {
		income := daysOfWeek[i]
		totalIncome += v
		fmt.Println(income, i)
	}

	fmt.Println("___________________")

	testsClice := []string{" test ", "tefst", "test", "test", "test", "te76st", "ytuest", "notest", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "tehgst", "test", "test", "test", "test"}
	var quantityTest int
	var quantityNoTest int

TestQuantity:
	for _, v := range testsClice {
		switch {
		case v == "test":
			quantityTest++
		case quantityTest >= 10:
			continue TestQuantity
		case quantityNoTest >= 3:
			fmt.Println("programm have a mistaces")
			break TestQuantity
		default:
			quantityNoTest++
		}

		fmt.Println(quantityNoTest, quantityTest)
	} /*<-вот тут конец TestQuantity*/

	questionVar := "i am in TestQuantity or not?"
	var answer string
	fmt.Println(questionVar)
	fmt.Scan(&answer)
	if answer == "not" || answer == "n" {
		fmt.Println("correct answer") /*коцец у TestQuantity там же где конец цикла фор*/
	}
}
