package main

import "fmt"

func main() {
	var age int

	fmt.Print("enter age: ")
	fmt.Scan(&age)

	fmt.Println("you", age, "years")
}
