package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 100

type numbers [size]int

func initialize(nums *numbers) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		nums[i] = rand.Intn(10000)
	}
}

func max(nums *numbers) int {
	temp := nums[0]
	for _, val := range nums {
		if val > temp {
			temp = val
		}
	}
	return temp
}

func main() {
	var nums *numbers = new(numbers)
	fmt.Println("before:", nums)
	initialize(nums)
	fmt.Println("after:", nums)
	fmt.Println("max:", max(nums))
}
