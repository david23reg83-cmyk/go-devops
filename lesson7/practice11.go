package main

import "fmt"

func main() {

	var number int

	fmt.Print("Введите число")
	fmt.Scan(&number)

	for i := 1; i <= number; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println("Другие значения", i)
		}
	}
}
