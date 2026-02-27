package main

import "fmt"

func main() {
	var number int

	fmt.Print("Ваше число: ")
	fmt.Scan(&number)

	if number > 100 {
		fmt.Print("Ваше число Больше 100")
	} else {
		fmt.Print("Ваше число подходит")
	}

}
