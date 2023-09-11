package main

import (
	"fmt"
	"math"
)

var (
	mul = func(op0, op1 int) int {
		return op0 * op1
	}

	sqr = func(val int) int {
		return mul(val, val)
	}
)

func main() {
	fmt.Printf("mul(25,7) = %d\n", mul(25, 7))
	fmt.Printf("sqr(13) = %d\n", sqr(13))

	fmt.Println("---")

	fmt.Printf(
		"94 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(94),
	)

	fmt.Println("---")

	for i := 0.0; i < 360.0; i += 45.0 {
		rad := func() float64 {
			return i * math.Pi / 180
		}()
		fmt.Printf("%.2f Deg = %.2f Rad\n", i, rad)
	}
}
