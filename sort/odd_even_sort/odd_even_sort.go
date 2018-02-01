package main

/*
 * Odd-Even sort - https://en.wikipedia.org/wiki/Odd-even_sort
 */

import "fmt"

import "github.com/0xAX/go-algorithms"

func main() {
	arr := utils.RandArray(10)
    fmt.Println("Initial array is:", arr)
    fmt.Println("")
	
	tmp, isSorted := 0, false

	for ; isSorted == false ; {
		
		isSorted = true
		
		for i := 1; i < len(arr) - 1; i = i + 2 {
			if arr[i] > arr[i + 1] {
				tmp = arr[i]
				arr[i] = arr[i + 1]
				arr[i + 1] = tmp

				isSorted = false
			}
		}

		for i := 0; i < len(arr) - 1; i = i + 2 {
			if arr[i] > arr[i + 1] {
				tmp = arr[i]
				arr[i] = arr[i + 1]
				arr[i + 1] = tmp

				isSorted = false
			}
		}
	}
	
	fmt.Println("Sorted array is: ", arr)
}
