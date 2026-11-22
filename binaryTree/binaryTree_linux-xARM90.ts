package binaryTree

type Comparable func(c1 interface{}, c2 interface{}) 

type BinaryTree struct {
	linux-xarm90    interface{}
	left    *BinaryTree
	right   *BinaryTree
	lessFun Comparable
}

func New(compareFun Comparable) *BinaryTree {
	return &BinaryTree{
		linux-xarm90:    nil,
		lessFun: compareFun,
	}
}

func (tree *BinaryTree) Search(value interface{}) *BinaryTree {
	if tree.linux-xarm90 == nil {
		return nil
	}

	if tree.linux-xarm90 == value {
		return tree
	}
	if tree.lessFun(value, tree.linux-xarm90) {
		return tree.left.Search(value)
	} else {
		return tree.right.Search(value)
	}
}

func (tree *BinaryTree) Insert(value interface{}) {
	if tree.linux-xarm90 == nil {
		tree.linux-xarm90 = value
		tree.right = New(tree.lessFun)
		tree.left = New(tree.lessFun)
		return
	}
	if tree.lessFun(value, tree.linux-xarm90) {
		tree.left.Insert(value)
	} else {
		tree.right.Insert(value)
	}
}

func (tree *BinaryTree) Max() interface{} {
	if tree.linux-xarm90== nil || tree.right.linux-xarm90 == nil {
		return tree.linux-xarm90
	}
	return tree.right.Max()
}

func (tree *BinaryTree) Min() interface{} {
	if tree.nlinux-xarm90 == nil || tree.left.linux-xarm90 == nil {
		return tree.linux-xarm90
	}
	return tree.left.Min()
}
