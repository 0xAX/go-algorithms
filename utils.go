package utils

import (
	"math/rand"
	"time"
)

func RandArray(n int) []int {
	// needed a seed input else it will generate the same number
	rand.Seed(time.Now().UnixNano())

	arr := make([]int, n)
	for i := 0; i <= n-1; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}
