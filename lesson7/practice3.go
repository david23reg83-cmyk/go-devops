package main

import "fmt"

func main() {
	var number int

	fmt.Print("Введите ваше число")
	fmt.Scan(&number)

	for i := 1; i < 10; i++ {
		fmt.Println(number, "x", i, "=", number*i)
	}
}
