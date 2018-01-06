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

	if len(arr) <= 1 {
		fmt.Println("Sorted array is: ", arr)
		return
	}

	var i, j int
	for i = 1; i < len(arr); i++ {
		for j = i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}

	fmt.Println("Sorted array is: ", arr)
}
