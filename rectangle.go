package main 

import "fmt"

func main() {
	var leng int
	var wid int

	square := leng * wid

	fmt.Print("Введите длину: ")
	fmt.Scan(&leng)

	fmt.Print("Введите ширину: ")
	fmt.Scan(&wid)

	fmt.Println("Площадь:", square)
}