package main

import (
	"fmt"
	"unsafe"
)

var (
	a uint8   = 72
	b int32   = 240
	c uint64  = 1234564321
	d float32 = 12432345.232
	e int64   = -1233453443434
	f float64 = -1.43555622362467
	g int16   = 32000
	h [5]rune = [5]rune{'O', 'n', 'T', 'o', 'p'}
)

func main() {
	fmt.Printf("a = %v [%T, %d bits]\n", a, a, unsafe.Sizeof(a)*8)
	fmt.Printf("b = %v [%T, %d bits]\n", b, b, unsafe.Sizeof(b)*8)
	fmt.Printf("c = %v [%T, %d bits]\n", c, c, unsafe.Sizeof(c)*8)
	fmt.Printf("d = %v [%T, %d bits]\n", d, d, unsafe.Sizeof(d)*8)
	fmt.Printf("e = %v [%T, %d bits]\n", e, e, unsafe.Sizeof(e)*8)
	fmt.Printf("f = %v [%T, %d bits]\n", f, f, unsafe.Sizeof(f)*8)
	fmt.Printf("g = %v [%T, %d bits]\n", g, g, unsafe.Sizeof(g)*8)
	fmt.Printf("h = %v [%T, %d bits]\n", h, h, unsafe.Sizeof(h)*8)

	fmt.Println("---")

	vals := []int{
		1024,
		0x0FF1CE,
		0x8BADF00D,
		0xBEEF,
		0777,
	}

	for _, i := range vals {
		if i == 0xBEEF {
			fmt.Printf("Got %d\n", i)
			break
		}
	}

	fmt.Println("---")

	p := 3.1415926535
	e := .5772156649
	x := 7.2e-5
	y := 1.616199e-35
	z := .416833e32

	fmt.Println(p, e, x, y, z)

	fmt.Println("---")

	a := -3.5 + 2i
	fmt.Printf("%v\n", a)
	fmt.Printf("%+g, %+g\n", real(a), imag(a))
}
