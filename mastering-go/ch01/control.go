package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide  command line argument")
		return
	}

	argument := os.Args[1]

	// With expression after switch
	switch argument {
	case "0":
		fmt.Println("Zero!")
	case "1":
		fmt.Println("One!")
	case "2", "3", "4":
		fmt.Println("2 or 3 or 4")
		fallthrough
	default:
		fmt.Println("value:", argument)
	}

	value, error := strconv.Atoi(argument)
	if error != nil {
		fmt.Println("Cannot convert o int:", argument)
		return
	}

	// No expression after switch
	switch {
	case value == 0:
		fmt.Println("Zero!")
	case value > 0:
		fmt.Println("Positive integer")
	case value < 0:
		fmt.Println("Negative integer")
	default:
		fmt.Println("This should not happen:", value)
	}
}
