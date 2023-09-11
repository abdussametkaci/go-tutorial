package main

import (
	"fmt"
)

var words = [][]string{
	{"break", "lake", "go", "right", "strong", "kite", "hello"},
	{"fix", "river", "stop", "left", "weak", "flight", "bye"},
	{"fix", "lake", "slow", "middle", "sturdy", "high", "hello"},
}

func search(w string) {
	isFound := false
DoSearch:
	for i := 0; i < len(words); i++ {
		for k := 0; k < len(words[i]); k++ {
			if words[i][k] == w {
				isFound = true
				break DoSearch
			}
		}
	}

	if isFound {
		fmt.Println("Found", w)
	} else {
		fmt.Println("Not Found", w)
	}
}

func search2(w string) {
DoSearch:
	for i := 0; i < len(words); i++ {
		for k := 0; k < len(words[i]); k++ {
			if words[i][k] == w {
				fmt.Println("Found", w)
				continue DoSearch
			}
		}
	}
}

func main() {
	search("left")
	search("asd")

	search2("fix")
}
