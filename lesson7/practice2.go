package main

import "fmt"

func main() {
	var number int

	fmt.Print("Укажите ваше число")
	fmt.Scan(&number)

	for i := 0; i < number; i++ {
		fmt.Println("Ваши числа:", i)
	}
}
