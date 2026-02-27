package main

import "fmt"

func main() {
	var age int
	var ticket int

	fmt.Print("Укажите ваш возраст: ")
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Print("есть ли членский билет: ")
		fmt.Scan(&ticket)

		if ticket == 1 {
			fmt.Println("Добро пожаловать")
		} else {
			fmt.Println("Нужен членский билет")
		}
	} else {
		fmt.Print("Вход запрещен")
	}
}
