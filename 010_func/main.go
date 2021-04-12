package main

import (
	"fmt"
)

// variadic function as it can take n inputs
func findAverage(nums ...int) float64 {
	num_len := len(nums)
	num_sum := 0
	for _, num := range nums {
		num_sum = num + num_sum
	}
	return float64(num_sum)/float64(num_len)
}

func main() {
	avg := findAverage(1,2,3,4, 7,9,165, 763)
	fmt.Println("Average of num is ", avg)
}