package stack

import "testing"

func Test_Stack(t *testing.T) {
	stack := New()

	stack.Push(5)
	stack.Push(6)
	stack.Push(7)

	if stack.Length() != 3 {
		t.Error("[Error] stack length is wrong")
	}

	stack.Pop()

	if stack.Length() != 2 {
		t.Error("[Error] stack length is wrong after pop")
	}

	if stack.Peek() != 6 {
		t.Error("[Error] stack Peek is wrong")
	}
}
