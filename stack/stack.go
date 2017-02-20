package stack

type StackItem struct {
	item interface{}
	next *StackItem
}

// Stack is a base structure for LIFO
type Stack struct {
	sp    *StackItem
	depth uint64
}

// Initialzes new Stack
func New() *Stack {
	var stack *Stack = new(Stack)

	stack.depth = 0
	return stack
}

// Pushes a given item into Stack
func (stack *Stack) Push(item interface{}) {
	stack.sp = &StackItem{item: item, next: stack.sp}
	stack.depth++
}

// Deletes top of a stack and return it
func (stack *Stack) Pop() interface{} {
	if stack.depth > 0 {
		item := stack.sp.item
		stack.sp = stack.sp.next
		stack.depth--
		return item
	}

	return nil
}

// Returns top of a stack without deletion
func (stack *Stack) Peek() interface{} {
	if stack.depth > 0 {
		return stack.sp.item
	}

	return nil
}
