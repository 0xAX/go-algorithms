package stack

import "testing"

func TestStack(t *testing.T) {
	var stack *Stack = New()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	for i := 5; i > 0; i-- {
		item := stack.Pop()

		if item != i {
			t.Error("TestStack failed...", i)
		}
	}
}
