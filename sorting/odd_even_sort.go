package main

/*
 * Odd-Even sort - https://en.wikipedia.org/wiki/Odd-even_sort
 */

func OddEvenSort(arr []int) {
	tmp, isSorted := 0, false

	for isSorted == false {

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
}
