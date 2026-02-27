package main

import "fmt"

func main() {
	var score int

	fmt.Print("Ваш колличество баллов: ")
	fmt.Scan(&score)

	if score >= 90 {
		fmt.Print("Отлично")
	} else if score >= 70 {
		fmt.Print("Хорошо")
	} else if score >= 50 {
		fmt.Print("Удовлетворително:")
	} else if score >= 30 {
		fmt.Print("Неудовлетворительно")
	} else {
		fmt.Print("Провал")
	}

}
