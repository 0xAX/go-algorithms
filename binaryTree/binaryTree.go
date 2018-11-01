package binaryTree

type Comparable func(c1 interface{}, c2 interface{}) bool

type BinaryTree struct {
	node    interface{}
	left    *BinaryTree
	right   *BinaryTree
	lessFun Comparable
}

func New(compareFun Comparable) *BinaryTree {
	return &BinaryTree{
		node:    nil,
		lessFun: compareFun,
	}
}

func (tree *BinaryTree) Search(value interface{}) *BinaryTree {
	if tree.node == nil {
		return nil
	}

	if tree.node == value {
		return tree
	}
	if tree.lessFun(value, tree.node) {
		return tree.left.Search(value)
	} else {
		return tree.right.Search(value)
	}
}

func (tree *BinaryTree) Insert(value interface{}) {
	if tree.node == nil {
		tree.node = value
		tree.right = New(tree.lessFun)
		tree.left = New(tree.lessFun)
		return
	}
	if tree.lessFun(value, tree.node) {
		tree.left.Insert(value)
	} else {
		tree.right.Insert(value)
	}
}

func (tree *BinaryTree) Max() interface{} {
	if tree.node == nil || tree.right.node == nil {
		return tree.node
	}
	return tree.right.Max()
}

func (tree *BinaryTree) Min() interface{} {
	if tree.node == nil || tree.left.node == nil {
		return tree.node
	}
	return tree.left.Min()
}
