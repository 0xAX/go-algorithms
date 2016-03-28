package queue

type Queue struct {
    queue []interface{}
    len int
}

func New() *Queue {
    queue := &Queue{}
    queue.queue = make([]interface{}, 0)
    queue.len = 0
    return queue
}

func (queue *Queue) Length() int {
    return queue.len
}

func (q *Queue) Remove() interface{} {
    tmp := q.queue[0]
    q.queue = q.queue[1:]
    q.len -= 1
    return tmp
}

func (q *Queue) Peek() interface{} {
    return q.queue[0]
}

func (q *Queue) Add(value interface{}) {
    q.len += 1
    q.queue = append(q.queue, value)
}
