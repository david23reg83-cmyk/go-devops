package main

import "fmt"

func main() {

	var number int
	var sum int

	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	for i := 1; i <= number; i++ {
		if i%2 == 0 {
			sum = sum + i
		}
	}
	fmt.Println("Ваши числа", sum)
}
