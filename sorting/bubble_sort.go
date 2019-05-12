package main

/*
 * Bubble sort - http://en.wikipedia.org/wiki/Bubble_sort
 */

import "fmt"

import "github.com/0xAX/go-algorithms"

func main() {
	arr := utils.RandArray(10)
	fmt.Println("Initial array is:", arr)
	fmt.Println("")

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}

	fmt.Println("Sorted array is: ", arr)
}
