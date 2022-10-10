package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Enter your name: ")
	var name string
	fmt.Scanln(&name)
	// fmt.Println("Your name is " + name)
	name = strings.TrimSpace(name) // removes spaces from start-end of a string
	fmt.Printf("Hi%s,I am Golang", name)
}
