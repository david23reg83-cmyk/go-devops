package main

import "fmt"

func main() {

	var age int
	var ticket int

	fmt.Print("Укажите ваш возраст: ")
	fmt.Scan(&age)

	fmt.Print("Есть ли членский билет: ")
	fmt.Scan(&ticket)

	if age >= 18 && ticket == 1 {
		fmt.Println("Добро пожаловать")
	} else if age >= 18 && ticket == 0 {
		fmt.Println("Нужен членский билет")
	} else {
		fmt.Println("Вход запрещен")
	}
}
