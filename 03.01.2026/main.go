package main

import (
    "fmt"
    "time"
)

func main(){
 var newMyYears = calcMyYears(1000)
 var yearMyBird, daysForMyDead = yearMyBird(newMyYears)
 yearMyBird = int(yearMyBird)
 fmt.Printf("My years: %.0f, мой год рождения: %d, до смерти мне осталось:%d дней\n",newMyYears, yearMyBird, daysForMyDead)
 testFunction("tester","my programm", 3)
 testFunction("programmer bog","world", 7)
 fmt.Println("________________________")
 var varForTestPointer int = 4
 
 testPointer1 := &varForTestPointer
 fmt.Printf("testPointer1\n")
 fmt.Printf("место в памяти переменной yearMyBird: %p или через &:%p,\n",testPointer1, &varForTestPointer)
 //будет ли оно так работать?
 *testPointer1 = 7
 // по идеи должно быть так: *varForTestPointer = 7
 
 /* & - компенсарт. используется для нахождения значение в памяти */
 //fmt.Printf("измененое значение varForTestPointer: %d через указатель testPointer1,\n",varForTestPointer)
 //fmt.Printf("место в памяти через testPointer1: %p и через комперсант varForTestPointer: %p\n", testPointer1, &varForTestPointer)
 fmt.Printf(`измененое значение varForTestPointer: %d через указатель testPointer1 место в памяти через testPointer1: %p и через комперсант varForTestPointer: %p`, varForTestPointer, testPointer1, &varForTestPointer)
 
 testPointer2 := 12520.00
 fmt.Println (testPointer2)
 testChangeVsrWithPointer (&testPointer2)
 fmt.Println (testPointer2)
}

func testChangeVsrWithPointer (test *float64){
 *test = *test * 26
}

func yearMyBird (myYears float64) (int, func(int)int) {
 myYears = myYears + 1 - 1
 currentYear := time.Now().Year()
 yearForMyBird := /*2026 хз как, но нужно научиться сюда подтягивать нынешнюю дату. Это делается так:*/ currentYear - int(myYears)
 howDaysForMyDead := func(myYears int)int{
  calculate := 100 - myYears
  calculate *= 364
  return calculate
 }
 
 return yearForMyBird, howDaysForMyDead
}

func calcMyYears (oldYears int) (newYears float64){
 newYears = float64 (1.00 + oldYears)
 return //naked return, голый возврат. можно, но не рекомендуется
}

func testFunction(name string, name2 string, intFerst int){
 aString := "i am"

 fmt.Printf("%s %s in %s %d days\n",aString, name, name2, intFerst)
}