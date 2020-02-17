package main

/*
 * Shell sort - http://en.wikipedia.org/wiki/Shellsort
 */
func ShellSort(arr []int) {
	for d := int(len(arr)/2); d > 0; d /= 2 {
		for i := d; i < len(arr); i++ {
			for j := i; j >= d && arr[j-d] > arr[j]; j -= d {
				arr[j], arr[j-d] = arr[j-d], arr[j]
			}
		}
	}
}
