package main

import "fmt"

func main() {

	var number int

	fmt.Println("Число:")
	fmt.Scan(&number)

	for i := 1; i <= 10; i++ {
		fmt.Println(number, "*", i, "=", number*i)
	}
}
