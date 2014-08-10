package binaryTree

import "testing"

func compare(x interface{}, y interface{}) bool {
    if x.(int) < y.(int) {
        return true
    } else {
        return false
    }
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
		t.Error("[Error] 2. Search erro")
	} 
}
