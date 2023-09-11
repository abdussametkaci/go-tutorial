package main

import (
	"fmt"
)

func main() {
	// Get user input
	fmt.Println("Please give me ypur name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Your name is", name)
}
