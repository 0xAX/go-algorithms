package main

import (
	"sync"
)

/* var list = []int{91, 28, 73, 46, 55, 64, 37, 82, 19}
var max = 1000

func main() {
	fmt.Println("before:", list)
	beadSort()
	fmt.Println("after: ", list)
} */

func BeadSort(list []int) {
	const bead = 'o'
	max := 1000

	all := make([]byte, max*len(list))
	abacus := make([][]byte, max)
	for pole, space := 0, all; pole < max; pole++ {
		abacus[pole] = space[:len(list)]
		space = space[len(list):]
	}
	var wg sync.WaitGroup
	wg.Add(len(list))
	for row, n := range list {
		go func(row, n int) {
			for pole := 0; pole < n; pole++ {
				abacus[pole][row] = bead
			}
			wg.Done()
		}(row, n)
	}
	wg.Wait()
	wg.Add(max)
	for _, pole := range abacus {
		go func(pole []byte) {
			top := 0
			for row, space := range pole {
				if space == bead {
					pole[row] = 0
					pole[top] = bead
					top++
				}
			}
			wg.Done()
		}(pole)
	}
	wg.Wait()
	for row := range list {
		x := 0
		for pole := 0; pole < max && abacus[pole][row] == bead; pole++ {
			x++
		}
		list[len(list)-1-row] = x
	}
}
