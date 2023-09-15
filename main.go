package main

import (
	"log"
	"math"
)

func main() {
	//nums := []int{2, -4, 2, -1, 3, -3, 10, -1, -11, -100, 8, -1} // 11
	//nums := []int{-4, -4, -2, -2, -3, -100} // -2
	nums := []int{-2, -3, 4, -1, -2, 1, 5, -3} // 7

	max := findConsecutiveSum(nums)

	log.Println(max)
}

func findConsecutiveSum(nums []int) int {
	// set max to minimum possible
	max := math.MaxInt * -1

	for i, _ := range nums {
		newMax := findMaxSumInArray(nums[i:])
		if newMax > max {
			max = newMax
		}
	}

	return max
}

func findMaxSumInArray(nums []int) int {
	max := math.MaxInt * -1
	sum := 0
	for _, n := range nums {
		sum += n
		if sum > max {
			max = sum
		}
	}

	return max
}
