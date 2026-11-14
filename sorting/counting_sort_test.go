package main

import "testing"

func TestCountingSortNegative(t *testing.T) {
	arr := []int{3, -1, 0, -5, 2, -1}

	CountingSort(arr)

	want := []int{-5, -1, -1, 0, 2, 3}
	for i := range want {
		if arr[i] != want[i] {
			t.Fatalf("got %v, want %v", arr, want)
		}
	}
}

func TestCountingSortEmpty(t *testing.T) {
	CountingSort([]int{})
}
