package list
 
import "sync"
 
type List struct {
    head *Item
    last *Item
    len int
    locker sync.RWMutex
}
 
type Item struct {
    Val interface{}
    next *Item
    prev *Item
    list *List
}
 
func New() *List {
    list := &List{}
    list.len = 0
    return list
}
 
func Insert(value interface{}, list *List) *List {
    newItem := &Item{value, list.head, list.last, list}
    list.locker.Lock()
    defer list.locker.Unlock()
 
    if list.head == nil {
        list.head = newItem
        list.last = newItem
    } else {
        list.head.prev = newItem
        list.head = newItem
        list.last.next = newItem
    }
 
    list.len++
 
    return list
}
 
func (list *List) First() *Item {
    return list.head
}
 
func (list *List) Last() *Item {
    return list.last
}
 
func (item *Item) Prev() *Item {
    return item.prev
}
 
func (item *Item) Next() *Item {
    return item.next
}
 
func Has(value interface{}, list *List) bool {
    if list.head == nil{
        return false
    }
    first := list.First()

    for {
        if first.Val == value {
            return true
        } else {
            if first.next != nil {
                first = first.next
            } else {
                return false
            }
        }
    }
 
    return false
}
 
func Remove(value interface{}, list *List) *List {
    list.locker.RLock()
    
    if list.head == nil {
    return list
    }

    list.locker.RUnlock()

    list.locker.RLock()
    first := list.First()
    last := list.Last()
    list.locker.RUnlock()
    list.locker.Lock()
    defer list.locker.Unlock()
    
    for {
        if last.next == nil {
            return list
        }
 
        if first.Val == value {
            first.prev.next = first.next
            first.next.prev = first.prev
            first.prev = nil
            first.next = nil
            first.Val = nil
            first.list = nil
            list.len--
            return list
        } else {
            first = first.next
        }
    }
}
 
func Length(list *List) int {
    return list.len
}
