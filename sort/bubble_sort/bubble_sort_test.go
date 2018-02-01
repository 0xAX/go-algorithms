package bubble_sort

import (
	sortable "github.com/0xAX/go-algorithms/sort"
	"math/rand"
	"testing"
)

func TestSort(t *testing.T) {
	var s = make(sortable.IntSlice, 10)

	for i := 0; i < 10; i++ {
		s[i] = rand.Intn(i + 1)
	}

	Sort(&s)

	for i := 0; i < 10; i++ {
		if i != 9 {
			if s[i] > s[i+1] {
				t.Fatal("s[i] > s[i + 1]", s[i], s[i+1])
			}
		}
	}
}
