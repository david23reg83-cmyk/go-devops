package main 

import "fmt"

func main() {
	var city string 
	var years int 
	
	fmt.Print("Введите город: ")
	fmt.Scan(&city)

	fmt.Print("Сколько лет живет: ")
	fmt.Scan(&years)

	fmt.Println("Ты живешь в", city, "уже", years, "лет")
}
