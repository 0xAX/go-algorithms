package binaryTree

import "testing"

func compare(x interface{}, y interface{}) bool {
	return x.(int) < y.(int)
}

func Test_binaryTree(t *testing.T) {
	tree := New(compare)

	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)

	findTree := tree.Search(2)
	if findTree.node != 2 {
		t.Error("[Error] Search error")
	}

	findNilTree := tree.Search(100)

	if findNilTree != nil {
		t.Error("[Error] 2. Search error")
	}
}

func Test_minmax(t *testing.T) {
	tree := New(compare)

	testValues := []int{4, 5, 3, 2, 9}
	for _, i := range testValues {
		tree.Insert(i)
	}

	max := tree.Max()
	if max != 9 {
		t.Errorf("[Error] max: expected 9, got %d", max)
	}

	min := tree.Min()
	if min != 2 {
		t.Errorf("[Error] max: expected 2, got %d", min)
	}
}
