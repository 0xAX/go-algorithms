package queue

import "testing"

func Test_Queue(t *testing.T) {
    queue := New()

    queue.Add(1)
    queue.Add(2)
    queue.Add(3)

    if queue.Length() != 3 {
        t.Error("[Error] queue length is wrong")
    }

    element := queue.Remove()

    if element != 1 {
        t.Error("[Error] remove is wrong")
    }

    if queue.Length() != 2 {
        t.Error("[Error] queue length is wrong after pop")
    }
    
    if queue.Peek() != 2 {
        t.Error("[Error] queue Peek is wrong")
    }
}
