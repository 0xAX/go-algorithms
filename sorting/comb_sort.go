package main

/*
 * Comb sort - https://en.wikipedia.org/wiki/Combsort
 */

import "fmt"

import "github.com/0xAX/go-algorithms"

func main() {
	arr := utils.RandArray(10)
	fmt.Println("Initial array is:", arr)
	fmt.Println("")

	tmp := 0
	arrLen := len(arr)
	gap := arrLen
	for gap > 1 {
		gap = gap * 10 / 13 //shrink factor is 1.3

		for i := 0; i+gap < arrLen; i++ {
			if arr[i] > arr[i+gap] {
				tmp = arr[i]
				arr[i] = arr[i+gap]
				arr[i+gap] = tmp
			}
		}
	}

	fmt.Println("Sorted array is: ", arr)
}
