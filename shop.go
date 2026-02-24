package main

import "fmt"

func main() {
	var price int
	var gaveMoney int

	fmt.Print("Цена: ")
	fmt.Scan(&price)

	fmt.Print("Дал денег: ")
	fmt.Scan(&gaveMoney)

	change := gaveMoney - price

	fmt.Println("Сдача:", change)

}
