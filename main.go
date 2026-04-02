package main

import "fmt"

func factorial(n int) uint64 {
	var chislo uint64 = 1
	for i := 2; i <= n; i++ {
		chislo *= uint64(i)
	}
	return chislo
}

func main() {
	var chisloFact int
	fmt.Print("Введите Ваше число: ")
	fmt.Scanln(&chisloFact)
	fmt.Printf("Факториал %d равен: %d\n", chisloFact, factorial(chisloFact))

}
