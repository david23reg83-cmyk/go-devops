package main

import "fmt"

func main() {

	var number int
	var count int

	fmt.Print("Введите число :")
	fmt.Scan(&number)

	for i := 1; i <= number; i++ {
		if i%2 == 0 {
			count = count + 1
		}
	}
	fmt.Println("Четных чилсел:", count)
}
