package main

import "fmt"

func main() {

	var val [100]int = [100]int{44, 72, 12, 55, 64, 1, 4, 90, 13, 54}
	var days [7]string = [7]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}
	var truth = [256]bool{true}
	var histogram = [5]map[string]int{
		{"A": 12, "B": 1, "D": 15},
		{"man": 1344, "women": 844, "children": 577},
	}

	/* // second way
	var histogram = [5]map[string]int {
	map[string]int{"A":12,"B":1, "D":15},
	map[string]int{"man":1344,"women":844, "children":577,...},
	}
	*/

	var board = [4][2]int{
		{33, 23},
		{62, 2},
		{23, 4},
		{51, 88},
	}
	var matrix = [2][2][2][2]byte{
		{{{4, 4}, {3, 5}}, {{55, 12}, {22, 4}}},
		{{{2, 2}, {7, 9}}, {{43, 0}, {88, 7}}},
	}

	var weekdays = [...]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
	}

	var msg = [12]rune{0: 'H', 2: 'E', 4: 'L', 6: 'O', 8: '!'}

	fmt.Println("val:", val)
	fmt.Println("days:", days)
	fmt.Println("truth:", truth)
	fmt.Println("histogram:", histogram)
	fmt.Println("board:", board)
	fmt.Println("matrix:", matrix)
	fmt.Println("weekdays:", weekdays)
	fmt.Println("msg:", msg)
}
