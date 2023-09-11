package main

import (
	"fmt"
	"strings"
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

func isDollar(curr Curr) bool {
	var result bool
	switch curr {
	default:
		result = false
	case Curr{"AUD", "Australian Dollar", "Australia", 36}:
		result = true
	case Curr{"HKD", "Hong Kong Dollar", "Hong Koong", 344}:
		result = true
	case Curr{"USD", "US Dollar", "United States", 840}:
		result = true
	}
	return result
}

func isDollar2(curr Curr) bool {
	dollars := []Curr{currencies[2], currencies[6], currencies[9]}
	switch curr {
	default:
		return false
	case dollars[0]:
		fallthrough
	case dollars[1]:
		fallthrough
	case dollars[2]:
		return true
	}
}

func isEuro(curr Curr) bool {
	switch curr {
	case currencies[2], currencies[4], currencies[10]:
		return true
	default:
		return false
	}
}

func find(name string) {
	isFound := false
	for i := 0; i <= 10; i++ {
		c := currencies[i]
		switch {
		case strings.Contains(c.Currency, name),
			strings.Contains(c.Name, name),
			strings.Contains(c.Country, name):
			fmt.Println("Found", c)
			isFound = true
		}
	}

	if !isFound {
		fmt.Println("Not Found", name)
	}
}

func assertEuro(c Curr) bool {
	switch name, curr := "Euro", "EUR"; {
	case c.Name == name:
		return true
	case c.Currency == curr:
		return true
	}
	return false
}

func findNumber(num int) {
	isFound := false
	for _, curr := range currencies {
		if curr.Number == num {
			fmt.Println("Found", curr)
			isFound = true
		}
	}

	if !isFound {
		fmt.Println("Not Found", num)
	}
}

func findAny(val interface{}) {
	switch i := val.(type) {
	case int:
		findNumber(i)
	case string:
		find(i)
	default:
		fmt.Printf("Unable to search with type %T\n", val)
	}
}

func main() {
	curr := Curr{"EUR", "Euro", "Italy", 978}
	if isDollar(curr) {
		fmt.Printf("%+v is Dollar currency\n", curr)
	} else if isEuro(curr) {
		fmt.Printf("%+v is Euro currency\n", curr)
	} else {
		fmt.Println("Currency is not Dollar or Euro")
	}
	dol := Curr{"HKD", "Hong Kong Dollar", "Hong Koong", 344}
	if isDollar2(dol) {
		fmt.Println("Dollar currency found:", dol)
	}

	find("EUR")

	if assertEuro(Curr{"EUR", "Euro", "Belgium", 978}) {
		fmt.Println("Yeap, it is euro")
	} else {
		fmt.Println("Nope, it is not euro")
	}

	findAny("Peso")
	findAny(404)
	findAny(978)
	findAny(false)
}
