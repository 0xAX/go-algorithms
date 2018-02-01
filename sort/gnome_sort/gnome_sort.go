package main

/*
 * Gnome sort - https://en.wikipedia.org/wiki/Gnome_sort
 */

import "fmt"

import "github.com/0xAX/go-algorithms"

func main() {
	arr := utils.RandArray(10)
    fmt.Println("Initial array is:", arr)
    fmt.Println("")
	
	i := 1
	tmp := 0
	for ; i < len(arr) ; {
		if arr[i] >= arr[i - 1] {
			i++
		} else {
			tmp = arr[i]
			arr[i] = arr[i - 1]
			arr[i - 1] = tmp

			if i > 1 {
				i--
			}
		}
	}

	fmt.Println("Sorted array is: ", arr)
}

