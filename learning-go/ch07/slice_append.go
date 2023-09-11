package main

import "fmt"

func main() {
	months := make([]string, 3, 3)
	months = append(months, "Jan", "Feb", "March",
		"Apr", "May", "June")
	months = append(months, []string{"Jul", "Aug", "Sep"}...)
	months = append(months, "Oct", "Nov", "Dec")
	fmt.Println(len(months), cap(months), months)
}
