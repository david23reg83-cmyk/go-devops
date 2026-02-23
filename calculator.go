package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	fmt.Print("Введите первое число: ")
	fmt.Scan(&num1)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&num2)

	sum := num1 + num2
	dif := num1 - num2
	work := num1 * num2
	divi := num1 / num2

	fmt.Println("Сумма:", sum)
	fmt.Println("Разность:", dif)
	fmt.Println("Произведение:", work)
	fmt.Println("Деление:", divi)
}
