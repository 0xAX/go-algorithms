package binaryTree

type Comparable func (c1 interface{}, c2 interface{}) bool

type BinaryTree struct {
    node interface{}
    left  *BinaryTree
    right *BinaryTree
    lessFun Comparable
}

func New(compareFun Comparable) *BinaryTree {
    tree := &BinaryTree{}
    tree.node = nil
    tree.lessFun = compareFun
    return tree
}

func (tree *BinaryTree) Search(value interface{}) *BinaryTree {
    if tree.node == nil {
        return nil
    }
    
    if tree.node == value {
        return tree
    } else {
        if tree.lessFun(value, tree.node) == true {
            t := tree.left.Search(value)
            return t
        } else {
            t := tree.right.Search(value)
            return t
        }
    }
}

func (tree *BinaryTree) Insert(value interface{}) {
    if tree.node == nil {
        tree.node = value
        tree.right = New(tree.lessFun)
        tree.left  = New(tree.lessFun)
        return
    } else {
        if tree.lessFun(value, tree.node) == true {
            tree.left.Insert(value)
        } else {
            tree.right.Insert(value)
        }
    }
}


