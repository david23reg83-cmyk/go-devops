package main

import "fmt"

func main() {
	var age int
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Println("Ты совершеннолетний")
	} else {
		fmt.Println("Ты не совершеннолетний")
	}
}
