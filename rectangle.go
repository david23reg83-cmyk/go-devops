package main

import "fmt"

func main() {
	var leng int
	var wid int

	fmt.Print("Введите длину: ")
	fmt.Scan(&leng)

	fmt.Print("Введите ширину: ")
	fmt.Scan(&wid)

	square := leng * wid

	fmt.Println("Площадь:", square)
}
