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
	array := []int{-2, -5, 6, 0, -2, 0, -3, 1, 0, 5, -6}
	fmt.Println("Maximum subarray sum: ", maxSubarray(array))
}
