package main

/*
 * Counting sort - https://en.wikipedia.org/wiki/Counting_sort
 */

// getRange returns the smallest value in arr and the number of distinct
// values between the smallest and the largest, inclusive.
func getRange(arr []int) (int, int) {
	if len(arr) == 0 {
		return 0, 1
	}

	min, max := arr[0], arr[0]

	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return min, max - min + 1
}

func CountingSort(arr []int) {
	min, k := getRange(arr)
	array_of_counts := make([]int, k)

	// Count each value at its offset from the minimum, so negative values
	// get a valid index too.
	for i := 0; i < len(arr); i++ {
		array_of_counts[arr[i]-min] += 1
	}

	for i, j := 0, 0; i < k; i++ {
		for {
			if array_of_counts[i] > 0 {
				arr[j] = i + min
				j += 1
				array_of_counts[i] -= 1
				continue
			}
			break
		}
	}
}
