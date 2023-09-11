package main

import "fmt"

func main() {
	var (
		txt2 = "\u6C34\x20brings\x20\x6c\x69\x66\x65."
		txt3 = ` 
		\u6C34\x20 
		brings\x20 
		\x6c\x69\x66\x65. 
		`
	)

	fmt.Print(txt2)
	fmt.Print(txt3)
}
