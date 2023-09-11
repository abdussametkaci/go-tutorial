package main

import "fmt"

func main() {
	var vector []float64
	fmt.Println(len(vector)) // prints 0, no panic
	fmt.Println(cap(vector)) // prints 0, no panic
	h := make([]float64, 4, 10)
	fmt.Println(len(h), ",", cap(h))
}
