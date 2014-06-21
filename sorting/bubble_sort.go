package main

/*
 * Bubble sort, sometimes referred to as sinking sort, 
   is a simple sorting algorithm that works by repeatedly 
   stepping through the list to be sorted, comparing 
   each pair of adjacent items and swapping them if 
   they are in the wrong order. The pass through 
   the list is repeated until no swaps are needed, 
   which indicates that the list is sorted.

   Complexity - О(n2)
 */

import "fmt"

import "github.com/IoSync/go-alghoritms"

func main() {
	arr := utils.RandArray(10)
	fmt.Println("Initial array is:", arr)
	fmt.Println("")
	
	tmp := 0
	
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr) - 1; j++ {
			if arr[j] > arr[j + 1] {
				tmp = arr[j]
				arr[j] = arr[j + 1]
				arr[j + 1] = tmp
			}
		}
	}
	
	fmt.Println("Sorted array is: ", arr)
}
