package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 4)
	ch <- 2
	ch <- 4
	ch <- 6
	ch <- 8
	// ch <- 10 // fatal error: all goroutines are asleep - deadlock!

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
