// кароче дело было к ночи(0:03) интерфес по сути контролер или тупая ии которая воспринимает только текст мы в нее фото не загрузим. интерфейс по сути просто посредник для того что бы выводы разных вункций прировнять к 1 виду
package main

import "fmt"

type secondBookName interface {
	SecondName() string
}

type secondNameBook struct {
	Author string
	Title  string
	Price  string
}

func (b secondNameBook) SecondName() string {
	return fmt.Sprintf("%s, автор: %s, price: %s", b.Title, b.Author, b.Price)
}

type nameBookShop struct {
	name string
}

func (n nameBookShop) SecondName() string {
	return fmt.Sprintf("shop for book: %s", n.name)

}

func main() {
	// Использование интерфейса
	var b secondBookName
	var s secondBookName

	b = secondNameBook{
		Author: "Антон",
		Title:  "Go для начинающих",
		Price:  "10",
	}

	s = nameBookShop{
		name: "testShop",
	}

	fmt.Println(b.SecondName(), s.SecondName())
}
