package main

import "fmt"

func main() {
	const amountSpend float64 = 10.53417282
	newLoaltyPoint := calculateNewPoint(amountSpend) /*можно вставлять как значения (10;1;8928 и тд, так и функции, константы */
	var loaltyPoint int = 14
	loaltyPoint = calculatePoint(newLoaltyPoint, loaltyPoint)
	fmt.Println(loaltyPoint, amountSpend)
}

func calculateNewPoint(amountSpend float64) int {
	loaltyPoint := int(2 * amountSpend)
	return loaltyPoint
}

func calculatePoint(newLoaltyPoint int, loaltyPoint int) int {
	calculatePointSumm := newLoaltyPoint + loaltyPoint
	return calculatePointSumm

}
