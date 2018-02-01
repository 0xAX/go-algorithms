package comb_sort

import (
	sortable "github.com/0xAX/go-algorithms/sort"
)

func Sort(data sortable.Sortable) sortable.Sortable {
	gap := data.Len()

	for {
		if gap > 1 {
			gap = gap * 100 / 124
		}

		for i := 0; ; {

			if data.Less(i+gap, i) {
				data.Swap(i, i+gap)
			}

			i++

			if i+gap >= data.Len() {
				break
			}
		}

		if gap == 1 {
			break
		}
	}

	return data
}
