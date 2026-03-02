package main

import "fmt"

func main() {

	var speed int

	fmt.Print("Укажите вашу скорость")
	fmt.Scan(&speed)

	if speed > 120 {
		fmt.Println("Скорость привышена")
	} else if speed > 60 {
		fmt.Println("средняя скорость")
	} else {
		fmt.Println("Медленно")
	}
}
