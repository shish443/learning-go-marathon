package main

import "fmt"

// BookName - интерфейс для сущностей, имеющих название
type BookName interface {
	Name() string
}

// Book - структура, описывающая книгу
type Book struct {
	Author string
	Title  string
}

// Name реализует интерфейс BookName
func (b Book) Name() string {
	return fmt.Sprintf("%s, автор: %s", b.Title, b.Author)
}

func main() {
	// Использование интерфейса
	var b BookName

	b = Book{
		Author: "Антон",
		Title:  "Go для начинающих",
	}

	fmt.Println(b.Name())
}
