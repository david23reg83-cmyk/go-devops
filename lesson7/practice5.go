package main

import "fmt"

func main() {

	var number int

	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	for i := 1; i <= number; i++ {
		fmt.Println(i, "* 2 =", i*2)
	}
}
