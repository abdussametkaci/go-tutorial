package main

import "fmt"

var valPtr *float32
var countPtr *int
var person *struct {
	name string
	age  int
}
var matrix *[1024]int
var row []*int64

func main() {
	fmt.Println(valPtr, countPtr, person, matrix, row)

	fmt.Println("---")

	var a int = 1024
	var aptr *int = &a

	fmt.Printf("a=%v\n", a)
	fmt.Printf("aptr=%v\n", aptr)

	fmt.Println("---")

	structPtr := &struct{ x, y int }{44, 55}
	pairPtr := &[2]string{"A", "B"}

	fmt.Printf("struct=%#v, type=%T\n", structPtr, structPtr)
	fmt.Printf("pairPtr=%#v, type=%T\n", pairPtr, pairPtr)
}
