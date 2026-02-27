package main

import "fmt"

func main() {

	var number1 int
	var number2 int
	var number3 int

	fmt.Print("Число1: ")
	fmt.Scan(&number1)

	fmt.Print("Число2: ")
	fmt.Scan(&number2)

	fmt.Print("Число3: ")
	fmt.Scan(&number3)

	// среднее арифметическое

	arith := (number1 + number2 + number3) / 3

	fmt.Println("Среднее:", arith)
}
