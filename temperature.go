package main

import "fmt"

func main() {
	var celsius int
	fmt.Print("Темепература в Цельсиях :")
	fmt.Scan(&celsius)

	fahrenheit := celsius*9/5 + 32

	fmt.Println("Фарингейт:", fahrenheit)
}
