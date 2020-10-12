package numerical

import (
	"fmt"
)

/* O(n) solution for calculating maximum subarray sum. */
func Max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxSubarray(array []int) int {
	var currentMax int = 0
	var maxTillNow int = 0
	for _, v := range array {
		currentMax = Max(v, currentMax+v)
		maxTillNow = Max(maxTillNow, currentMax)
	}
	return maxTillNow
}

func main() {
	array := []int{-3, -4, 7, 1, -2, 0, -5, 1, 0, 6, -5}
	fmt.Println("Maximum subarray sum: ", maxSubarray(array))

	array = []int{3, 4, -7, 2, 0, 0, -3, -1, 0, -5, 7}
	fmt.Println("Maximum subarray sum: ", maxSubarray(array))
}
