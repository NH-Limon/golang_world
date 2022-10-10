package main

import "fmt"

func main() {
	fmt.Println("Welcome to my madlib game")
	fmt.Print("Enter your favourite color: ")
	var color string
	fmt.Scanln(&color)
	fmt.Print("Enter your favourite hobby: ")
	var hobby string
	fmt.Scanln(&hobby)
	fmt.Print("Enter your address: ")
	var address string
	fmt.Scanln(&address)
	fmt.Printf("Your favourite color is: %s, favourite hobby is: %s and address is: %s.\n", color, hobby, address)
}
