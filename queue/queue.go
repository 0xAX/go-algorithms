/* queue.go - A generic queue implementation
 *
 * This file is part of ds library.
 *
 * The MIT License (MIT)
 * Copyright (c) <2016> Alexander Kuleshov <kuleshovmail@gmail.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
 * of the Software, and to permit persons to whom the Software is furnished to do
 * so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
 * INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A
 * PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
 * HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF
 * CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
 * OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */
 
package queue

type QueueItem struct {
	item interface{}
	prev *QueueItem
}

// Base data structure for Queue
type Queue struct {
	current *QueueItem
	last *QueueItem
	depth uint64
}

// Initializes new Queue and return it
func New() *Queue {
	var queue *Queue = new(Queue)

	queue.depth = 0

	return queue
}

// Puts a given item into Queue
func (queue *Queue) Enqueue(item interface{}) {
	if (queue.depth == 0) {
		queue.current = &QueueItem{item: item, prev: nil}
		queue.last = queue.current
		queue.depth++
		return
	}
		
	q := &QueueItem{item: item, prev: nil}
	queue.last.prev = q
	queue.last = q
	queue.depth++
}

// Extracts first item from the Queue
func (queue *Queue) Dequeue() interface{} {
	if (queue.depth > 0) {
		item := queue.current.item
		queue.current = queue.current.prev
		queue.depth--
		
		return item
	}

	return nil
}
