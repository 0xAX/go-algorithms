package cocktal_sort

import (
	sortable "github.com/0xAX/go-algorithms/sort"
)

func Sort(data sortable.Sortable) sortable.Sortable {
	for i := 0; i < data.Len()/2; i++ {
		left := 0
		right := data.Len() - 1

		for left <= right {
			if data.Less(left+1, left) {
				data.Swap(left, left+1)
			}

			left++

			if data.Less(right, right-1) {
				data.Swap(right, right-1)
			}

			right--
		}
	}

	return data
}
