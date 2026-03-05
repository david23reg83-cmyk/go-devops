package main

import "fmt"

func main() {

	var age int
	var ticket int

	fmt.Print("Укажите ваш возраст: ")
	fmt.Scan(&age)

	fmt.Print("укажите билет: ")
	fmt.Scan(&ticket)

	if age >= 18 && ticket == 1 {
		fmt.Println("Проходите")
	} else {
		fmt.Println("Вход запрещен")
	}

}
