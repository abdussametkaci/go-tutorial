package main

import (
	"fmt"
)

type shape interface {
	area() float64
}

type polygon interface {
	area() float64
}

type rect struct {
	name           string
	length, height float64
}

func (r *rect) area() float64 {
	return r.length * r.height
}

func (r *rect) String() string {
	return fmt.Sprintf(
		"%s[lenght:%.2f height:%.2f]",
		r.name, r.length, r.height,
	)
}

type triangle struct {
	name    string
	a, b, c float64
}

func (t *triangle) area() float64 {
	return 0.5 * (t.a * t.b)
}

func (t *triangle) String() string {
	return fmt.Sprintf(
		"%s[sides: a=%.2f b=%.2f c=%.2f]",
		t.name, t.a, t.b, t.c,
	)
}

func shapeInfo(s shape) string {
	return fmt.Sprintf("Shape Area = %.2f", s.area())
}

func polygonInfo(s polygon) string {
	return fmt.Sprintf("Polygon Area = %.2f", s.area())
}

func main() {
	r := &rect{"shape Square", 4.0, 4.0}
	fmt.Println(r, "=>", shapeInfo(r))

	t := &triangle{"shape Right Triangle", 1, 2, 3}
	fmt.Println(t, "=>", shapeInfo(t))

	r2 := &rect{"polygon Square", 6.0, 6.0}
	fmt.Println(r2, "=>", polygonInfo(r2))

	t2 := &triangle{"polygon Right Triangle", 2, 3, 4}
	fmt.Println(t2, "=>", polygonInfo(t2))
}
