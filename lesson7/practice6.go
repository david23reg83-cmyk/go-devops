package main

import "fmt"

func main() {

	var number int

	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	for i := 1; i <= number; i++ {
		if i%2 == 0 {
			fmt.Println(number, "X", i, "=", number*i)
		}
	}
}
