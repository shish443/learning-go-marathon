package main

import (
	"fmt"
)

/*
как делать не следует:

	type BookShop struct {
		Name  string
		Greet func(shop BookShop)
	} как следует:
*/
type BookShop struct {
	Name string
}

type Ooo int
type TestStruct struct {
	TestMeaningString  string
	TestMeaningInt     int
	TestMeaningBool    bool
	TestMeaningFloat64 float64
} /*структура struct фиксированный, его нельзя изменить, он как const или array, сами значения structa изменять можно*/

type StructTest struct {
	Test1 int
	Test2 int
	Test3 int
}

func main() {
	/* как делать не следует 3:
	myShop := BookShop{
		Name:  "reven's shop",
		Greet: greetShop,
	}
	myShop.Greet(myShop)
	как следует:*/
	myShop := BookShop{
		Name: "reven's shop",
	}
	myShop.Greet()

	fmt.Println("_______________________________________________")

	var test TestStruct

	fmt.Println("что тут?", test) /*нулевые значения для всего struct*/

	test.TestMeaningBool = true
	test.TestMeaningFloat64 = 435.4532523
	test.TestMeaningInt = 21
	test.TestMeaningString = "something"

	fmt.Println("что теперь тут?", test)

	fmt.Println("_______________________________________________")

	test2 := TestStruct{ /*struct literal*/
		TestMeaningFloat64: 23.23425, /*в данном случае порядок следования не важен*/
		TestMeaningInt:     12,
		TestMeaningString:  "something",
		TestMeaningBool:    true,
	} /*можно указать значениятолько для части полей*/
	// или так краткий вариант:

	fmt.Printf("что в test2? %v %p \n", test2, &test2) /*места в памяти для этого стракта и для стракта в тестзаменастракт будут разными*/

	test2 = TestStruct{"something", 12, true, 23.23425} /*в данном случае порядок следования важен*/
	testTransferStruct(test2)                           /*сдесь в этом стракте значение не поменяется*/

	fmt.Println("_______________________________________________")

	fmt.Printf("test2 до изменения по указателю %v \n", test2)
	pointerToTestMeaningInt := &test2
	testChangeStruct(pointerToTestMeaningInt)
	fmt.Printf("test2 после изменения по указателю %v \n", test2)

	fmt.Println("_______________________________________________")

	var myTest Ooo = 12345
	myTest.TestInt()

	fmt.Println("_______________________________________________")
	MapVsStruct()
	fmt.Println("_______________________________________________")

}

func testTransferStruct(testTransferStruct TestStruct) {
	fmt.Printf("стракт до изменения %v %p \n", testTransferStruct, &testTransferStruct)

	testTransferStruct.TestMeaningBool = false

	fmt.Printf("стракт после изменения %v %p \n", testTransferStruct, &testTransferStruct)
	/*стракт копируется, а не передается указателем, как саммив, но не как мап*/

	fmt.Println("_______________________________________________")
}

func testChangeStruct(TestMeaningInt *TestStruct) {
	TestMeaningInt.TestMeaningInt = 14
}

/*как делать не следует 2:
func greetShop(shop BookShop) {
	fmt.Println("welcome to", shop.Name)
}как следует:*/

func (shop BookShop) Greet() {
	fmt.Println("welcome to", shop.Name)
}

func (test Ooo) TestInt() {
	fmt.Println(test)
}

func MapVsStruct() {
	// Struct version
	testStruct := StructTest{
		Test1: 1,
		Test2: 2,
		Test3: 3,
	}

	// Map version
	testMap := map[string]int{
		"Test1": 1,
		"Test2": 2,
		"Test3": 3,
	}

	fmt.Println("---- Using Struct ----")
	fmt.Println("Test1:", testStruct.Test1) //используем .(точку), так как стракт уже на этапе компиляции финксированный
	fmt.Println("Test2:", testStruct.Test2)
	fmt.Println("Test3:", testStruct.Test3)

	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")

	fmt.Println("---- Using Map ----")
	fmt.Println("Test1:", testMap["CustomerName"]) //использует [], потому что мап изменяемый и на этапе компиляции кода компилятор не знает что именно внутри мапы
	fmt.Println("Test2:", testMap["CoffeeType"])
	fmt.Println("Test3:", testMap["CoffeeSize"])

}
