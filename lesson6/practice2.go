package main

import "fmt"

func main() {

	var discount int
	fmt.Print("Введите сумму покупки: ")
	fmt.Scan(&discount)

	if discount >= 10000 {
		fmt.Print("Скидка действует")
	} else {
		fmt.Print("Скидки нет")
	}

}
