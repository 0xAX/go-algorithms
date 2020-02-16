package main

import (
	"sort"
	"testing"

	utils "github.com/0xAX/go-algorithms"
)

var funcs = []struct{
	name string
	f	 Sort
}{
	{"shell", ShellSort},
	{"selection", SelectionSort},
	{"oddeven", OddEvenSort},
	{"insertion", InsertionSort},
	{"heap", HeapSort},
	{"gnome", GnomeSort},
	{"counting", CountingSort},
	{"comb", CombSort},
	{"cocktail", CocktailSort},
	{"bubble", BubbleSort},
}

func TestSort(t *testing.T) {
	for _, tt := range funcs {
		t.Run(tt.name, func(t *testing.T) {
			arr := utils.RandArray(10)

			tt.f(arr)

			if !sort.IntsAreSorted(arr) {
				t.Errorf("%v is not sorted", arr)
			}
		})
	}
}

func BenchmarkSort(b *testing.B) {
	for _, tt := range funcs {
		arrs := make([][]int, 0)

		for n := 0; n < b.N; n++ {
			arrs = append(arrs, utils.RandArray(10))
		}

		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < len(arrs); n++ {
				tt.f(arrs[n])
			}
		})
	}
}
