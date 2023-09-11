package main

import "fmt"

func main() {
	p := [5]int{122, 6, 23, 44, 6}
	p[4] = 82
	fmt.Println("p:", p)

	seven := [7]string{"grumpy", "sleepy", "bashful"}
	fmt.Println("len:", len(seven), "cap:", cap(seven))
}
