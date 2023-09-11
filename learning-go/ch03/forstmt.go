package main

import (
	"fmt"
	"math/rand"
)

type Curr struct {
	Currency string
	Name     string
	Country  string
	Number   int
}

var currencies = []Curr{
	Curr{"DZD", "Algerian Dinar", "Algeria", 12},
	Curr{"AUD", "Australian Dollar", "Australia", 36},
	Curr{"EUR", "Euro", "Belgium", 978},
	Curr{"CLP", "Chilean Peso", "Chile", 152},
	Curr{"EUR", "Euro", "Greece", 979},
	Curr{"HTG", "Gourde", "Haiti", 332},
	Curr{"HKD", "Hong Kong Dollar", "Hong Koong", 344},
	Curr{"USD", "US Dollar", "United States", 840},
	Curr{"USD", "US Dollar", "United States", 841},
	Curr{"USD", "US Dollar", "United States", 830},
	Curr{"EUR", "Euro", "ASD", 842},
}

var list1 = []string{
	"break", "lake", "go",
	"right", "strong",
	"kite", "hello"}

var list2 = []string{
	"fix", "river", "stop",
	"left", "weak", "flight",
	"bye"}

func listCurrs() {
	i := 0
	for i < len(currencies) {
		fmt.Println(currencies[i])
		i++
	}
}

func nextPair() (w1, w2 string) {
	pos := rand.Intn(len(list1))
	return list1[pos], list2[pos]
}

func printCurrencies() {
	for i := range currencies {
		fmt.Printf("%d: %v\n", i, currencies[i])
	}
}

func main() {
	listCurrs()
	fmt.Println("---")

	rand.Seed(31)

	for w1, w2 := nextPair(); w1 != "go" && w2 != "stop"; w1, w2 = nextPair() {
		fmt.Printf("Word Pair -> [%s, %s]\n", w1, w2)
	}

	fmt.Println("---")

	vals := []int{4, 2, 6}

	for i, v := range vals {
		// v--
		vals[i] = v - 1
	}

	fmt.Println(vals)

	fmt.Println("---")

	printCurrencies()

	fmt.Println("---")

	for range []int{1, 1, 1, 1} {
		fmt.Println("Looping")
	}
}
