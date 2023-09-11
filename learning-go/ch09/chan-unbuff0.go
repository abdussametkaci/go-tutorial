package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 12          // blocks
	fmt.Println(<-ch) // fatal error: all goroutines are asleep - deadlock!
}
