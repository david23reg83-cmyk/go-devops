package main

import "fmt"

func main() {
	var number int

	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	for i := 1; i <= number; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println("Ваши числа:", i)
		}
	}
}
