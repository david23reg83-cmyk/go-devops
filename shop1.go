package main

import "fmt"

func main() {
	var price1 int
	var price2 int
	var price3 int
	var money int

	// запрашиваем цены

	fmt.Print("Цена1: ")
	fmt.Scan(&price1)

	fmt.Print("цена2: ")
	fmt.Scan(&price2)

	fmt.Print("цена3: ")
	fmt.Scan(price3)

	//Спрашиваем, сколько денег

	fmt.Print("Сколько денег дал покупатель: ")
	fmt.Scan(&money)

	sum := price1 + price2 + price3

	change := money - sum

	fmt.Println("Цена1:", price1, "Цена2:", price2, "Цена3:", price3, "Денег получено:", money, "Сдача:", change)
}
