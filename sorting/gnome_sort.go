package main

/*
 * Gnome sort - https://en.wikipedia.org/wiki/Gnome_sort
 */
func GnomeSort(arr []int) {
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
}
