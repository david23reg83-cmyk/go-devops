package main

import "fmt"

func main() {
	var pin int
	var balance int
	var conclusion int

	fmt.Print("Введите пин-код: ")
	fmt.Scan(&pin)

	if pin == 3345 {
		fmt.Print("Введите ваш баланс: ")
		fmt.Scan(&balance)

		fmt.Print("Сколько вывести: ")
		fmt.Scan(&conclusion)

		if balance >= conclusion {
			fmt.Println("Деньги выданы")
		} else {
			fmt.Println("Недостаточно средств")
		}
	} else {
		fmt.Println("Неверный пин-код")
	}

}
