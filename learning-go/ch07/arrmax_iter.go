package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 100

var nums [size]int

func init() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		nums[i] = rand.Intn(10000)
	}
}

func max(nums [size]int) int {
	temp := nums[0]
	for _, val := range nums {
		if val > temp {
			temp = val
		}
	}
	return temp
}

func main() {
	fmt.Println("nums:", nums)
	fmt.Println("max:", max(nums))
}
