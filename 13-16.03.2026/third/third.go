package main

import "fmt"

type bookRaiting interface {
	Raiting() int
}

type firstMethod struct {
	UsersStars   int
	CriticsStars int
}

type secondMethod struct {
	UsersStars   int
	CriticsStars int
}

func (f firstMethod) Raiting() int {
	firstInt := f.UsersStars * f.CriticsStars
	return firstInt
}
func (s secondMethod) Raiting() int {
	secondInt := s.UsersStars + s.CriticsStars
	return secondInt
}
func PrintRaiting(r bookRaiting) {
	fmt.Println("райтинг это книги:", r.Raiting())
}

func main() {
	first := firstMethod{
		UsersStars:   2,
		CriticsStars: 13,
	}
	var second bookRaiting = secondMethod{
		UsersStars:   4,
		CriticsStars: 13,
	}
	PrintRaiting(first)
	PrintRaiting(second)

	second = firstMethod{
		UsersStars:   2,
		CriticsStars: 13,
	}
	PrintRaiting(second)

}
