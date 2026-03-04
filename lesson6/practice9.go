package main

import "fmt"

func main() {

	var temp int
	var rain int

	fmt.Print("Введите температуру: ")
	fmt.Scan(&temp)

	fmt.Print("Идет ли дождь: ")
	fmt.Scan(&rain)

	if temp < 0 || rain == 1 {

		fmt.Println("Возьми куртку:")
	} else {
		fmt.Println("Хорошая погода")

	}
}
