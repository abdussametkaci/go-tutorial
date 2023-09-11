package main

import "fmt"

func scale(factor float64, vector []float64) []float64 {
	for i := range vector {
		vector[i] *= factor
	}
	return vector
}

func contains(val float64, numbers []float64) bool {
	for _, num := range numbers {
		if num == val {
			return true
		}
	}
	return false
}

func main() {
	h := []float64{12.5, 18.4, 7.0}
	h[0] = 15
	fmt.Println("h[0]:", h[0])

	h10 := scale(2.0, h)
	fmt.Println("h10:", h10)
}
