package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter your birth year: ")
	var year int
	fmt.Scanln(&year)
	year = 2022 - year
	fmt.Print("Your age is: ")
	fmt.Print(year)
	fmt.Println()
}
