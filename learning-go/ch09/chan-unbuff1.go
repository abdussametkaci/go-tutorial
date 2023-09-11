package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() { ch <- 12 }()
	fmt.Println(<-ch)
}
