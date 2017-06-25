package main

/*
 * Insertion sort - https://en.wikipedia.org/wiki/Insertion_sort
 */

import "fmt"
import "github.com/0xAX/go-algorithms"

func main() {
	arr := utils.RandArray(10)
	fmt.Println("Initial array is:", arr)
	fmt.Println("")

	if (len(arr) <= 1) {
		fmt.Println("Sorted array is: ", arr)
		return
	}

	for curr := 1; curr < len(arr); curr++ {
		curr_val := arr[curr]
		prev := curr - 1
		for ; prev >= 0 && arr[prev] > curr_val; prev-- {
			arr[prev + 1] = arr[prev]
		}
		arr[prev + 1] = curr_val
	}

	fmt.Println("Sorted array is: ", arr)
}
