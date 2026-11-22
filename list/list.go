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
    list.locker.Lock()
    defer list.locker.Unlock()
 
    newItem := &Item{Val: value, list: list}
 
    if list.head == nil {
        list.head = newItem
        list.last = newItem
    } else {
        newItem.next = list.head
        list.head.prev = newItem
        list.head = newItem
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
    list.locker.Lock()
    defer list.locker.Unlock()

    for item := list.head; item != nil; item = item.next {
        if item.Val != value {
            continue
        }

        if item.prev == nil {
            list.head = item.next
        } else {
            item.prev.next = item.next
        }

        if item.next == nil {
            list.last = item.prev
        } else {
            item.next.prev = item.prev
        }

        item.prev = nil
        item.next = nil
        item.Val = nil
        item.list = nil
        list.len--
        return list
    }

    return list
}
 
func Length(list *List) int {
    return list.len
}
