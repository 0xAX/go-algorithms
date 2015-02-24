package main

import "fmt"

func main() {
	searchValue := 0

	arr := [10]int{1, 5, 100, 0, -100, 15, 4, 102, 30, 1000}
	fmt.Println(arr)

	// Sort the numbers
	tmp := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}

	fmt.Println(arr)

	left := 0
	right := len(arr) - 1

	if right < left {
		fmt.Println("Not found")
		return
	}

	// Find the number by looking at the center of the array, choosing
	// the left or right side depending on the value and then continue
	// to halve until the result has been found.
	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == searchValue {
			fmt.Println("Found at position: ", mid)
			return
		} else if arr[mid] < searchValue {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Println("Not found")
}
