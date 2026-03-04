package main

import "fmt"

func main() {

	var age int
	var sum int

	fmt.Print("Укажите ваш возраст: ")
	fmt.Scan(&age)

	fmt.Print("Укажите сумму: ")
	fmt.Scan(&sum)

	if age >= 18 && sum >= 10000 {
		fmt.Println("Можете купить автомобиль")
	} else if age >= 18 && sum <= 10000 {
		fmt.Println("Недостаточно денег")
	} else if age < 18 {
		fmt.Println("Тебе нельзя водить")
	}

}
