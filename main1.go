package main

import "fmt"

const pricePerKm float64 = 8.3   // руб. за км.
const pricePerMin float64 = 12.5 // руб. в мин.

func calculations(distance, time uint64) (float64, float64) {
	sumDistance := float64(distance) * pricePerKm
	sumTime := float64(time) * pricePerMin
	return sumDistance, sumTime
}

func main() {
	var distance, time uint64

	fmt.Print("Введите расстояние (в км.): ")
	fmt.Scanln(&distance)
	fmt.Print("Введите время (в мин.): ")
	fmt.Scanln(&time)

	sumDistance, sumTime := calculations(distance, time)
	fmt.Printf("Цена поездки (за км.): %.2f\n", sumDistance)
	fmt.Printf("Цена поездки (за мин.): %.2f\n", sumTime)

	var result string
	if sumDistance > sumTime {
		result = "по минутам"
	} else {
		result = "за киллометры"
	}
	fmt.Printf("Лучше использовать тариф %v\n", result)
}
