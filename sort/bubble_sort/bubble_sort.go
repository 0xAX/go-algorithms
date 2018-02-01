package bubble_sort

import (
	sortable "github.com/0xAX/go-algorithms/sort"
)

func Sort(data sortable.Sortable) sortable.Sortable {
	for i := 0; i < data.Len()-1; i++ {
		for j := 0; j < data.Len()-i-1; j++ {
			if data.Less(j+1, j) {
				data.Swap(j, j+1)
			}
		}
	}

	return data
}
