//ох как я соскучился любимый мой го, неболбшой перерыв на 2 месяца из-за экзаменовв, но я снова тут
// что же мне тут ИИшка набрасала план по всем темам, которые я прошел в 13 и 14 разделах, так что я пожалуй погнал писать код!!!!!
package main

import (
	"errors"
	"fmt"
)

// ЗАДАЧА 1: Создай интерфейс BookSeller.
// В нём должен быть один метод: SellBook(title string, amount int) (string, error)
type BookSeller interface{
	SellBook(title string, amount int) (string, error)
}

// ЗАДАЧА 2: Создай структуру Clerk (Продавец).
// Добавь в неё одно поле Name типа string.
type Clerk struct{
	Name string
}

// ЗАДАЧА 3: Реализуй метод SellBook для типа Clerk (чтобы структура удовлетворяла интерфейсу BookSeller).
// - Если количество книг (amount) <= 0, возвращай ошибку (через errors.New) с текстом "нельзя продать такое количество книг".
// - Если amount > 0, возвращай строку формата "Продавец [Имя] продал [amount] шт. книги '[title]'" и nil в качестве ошибки.

 func (c Clerk) SellBook(title string, amount int) (string, error) {
	if amount <= 0{
		return "", errors.New("нельзя продать такое количество книг")
	} else {
		return fmt.Sprintf("продавец %s продал %d шт. книг '%s' ", c.Name, amount, title), nil
	} 
 }

// ЗАДАЧА 4: Реализуй встроенный интерфейс Stringer для типа Clerk.
// Метод String() string должен возвращать красиво оформленную строку, например: "Сотрудник: [Имя]".

func (c Clerk) String() string{
	return fmt. Sprintf("Сотрудник: %s", c.Name)
}

// ЗАДАЧА 5: ПУСТОЙ ИНТЕРФЕЙС. Напиши функцию InspectItem(item interface{}).
// (Можешь использовать 'item any' - это одно и то же).
// Внутри используй fmt.Printf с глаголами %T (для типа) и %v (для значения), 
// чтобы вывести информацию о том, что именно передали в функцию.
// Формат вывода: "Тип: %T, Значение: %v\n"

func InspectItem(item interface{} /* /any */){
	fmt.Printf("тип-%T, значение-%v \n", item, item)
}

// ЗАДАЧА 6: Напиши функцию ProcessOrder(seller BookSeller, title string, amount int).
// 1. Сделай отложенный вызов (defer) анонимной функции.
// 2. Внутри анонимной функции используй recover(). Если была паника, выведи сообщение: "Отмена операции: <текст ошибки из паники>".
// 3. Вызови метод SellBook у переданного продавца (seller).
// 4. Если метод вернул ошибку (err != nil), принудительно вызови панику: panic(err).
// 5. Если ошибки нет, просто выведи строку с результатом продажи.

func ProcessOrder(seller BookSeller, title string, amount int){
	defer func(){
		if r := recover(); r!=nil{
			fmt.Println("Отмена операции:",r)
		}
	}()
	str,err:= seller.SellBook(title, amount)
	if err!=nil{
		panic(err)
	}
	fmt.Println(str)
}

func main() {
	// ЗАДАЧА 7: Создай экземпляр продавца (Clerk) с любым именем.
	shasha := Clerk{Name: "shasha"}

	fmt.Println("--- Проверка пустого интерфейса ---")
	// ЗАДАЧА 8: Передай в функцию InspectItem разные типы данных по очереди:
	// 1. Обычное число (например, 42)
	// 2. Строку (например, "Гарри Поттер")
	// 3. Своего продавца (Clerk)
	// Обрати внимание: при выводе продавца функция Printf сама найдет и использует твой метод String() из Задачи 4!

	InspectItem(42)
	InspectItem("Гарри Поттер")
	InspectItem(shasha)


	fmt.Println("\n--- Успешный заказ ---")
	// ЗАДАЧА 9: Вызови ProcessOrder, передав продавца, название любой книги и корректное количество (например, 3).

	ProcessOrder(shasha, "Загадай любовь", 3)

	fmt.Println("\n--- Проблемный заказ ---")
	// ЗАДАЧА 10: Вызови ProcessOrder с отрицательным или нулевым количеством книг.
	// Функция SellBook вернет ошибку, ProcessOrder вызовет панику, но defer с recover ее поймают.
	
	ProcessOrder(shasha, "Загадай любовь", -1)

	fmt.Println("\nМагазин продолжает работу без сбоев!")
}
