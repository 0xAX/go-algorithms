package queue

import "testing"

func TestQueue(t *testing.T) {
	nodes := []int{
		1, 2, 3, 4, 5, 6,
	}
	edges := [][]bool{
		{false, true, true, false, false, false},
		{true, false, false, true, false, false},
		{true, false, false, true, false, false},
		{false, true, true, false, true, false},
		{false, false, false, true, false, true},
		{false, false, false, false, true, false},
	}
	start := 1
	end := 6
	result := breadthFirstSearch(start, end, nodes, edges)
	expectedResult := []int{1, 3, 4, 5, 6}
	assert.Equal(t, result, expectedResult, "value should be equal")
}

