package stack

type Stack struct {
    st []interface{}
    len int
}

func New() *Stack {
    stack := &Stack{}
    stack.st = make([]interface{}, 1)
    stack.len = 0
    return stack
}

func (stack *Stack) Length() int {
    return stack.len
}

func (stack *Stack) Pop() {
    stack.st = stack.st[1:]
    stack.len -= 1
}

func (stack *Stack) Peek() interface{} {
    return stack.st[0]
}

func (stack *Stack) IsEmpty() bool {
    return (stack.len == 0)
}

func (stack *Stack) Push(value interface{}) {
    add(stack, value)
} 

func add(slice *Stack, value interface{}) {
    slice.len += 1
    var tmpSlice []interface{} = make([]interface{}, slice.len)
    if slice.len == 0 {
        slice.st[0] = value
        return
    }
    
    for i:=0; i < slice.len; i++ {
        tmpSlice[i] = 0
    }
    
    for i:=0; i < slice.len; i++ {
        if i == 0 {
            tmpSlice[0] = value
        } else {
            tmpSlice[i] = slice.st[i - 1]
        }

        if i == slice.len - 1 {
            break
        }
    }
    
    slice.st = tmpSlice
}
