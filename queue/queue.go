/*
Queue

- Collection of entities that are maintained in a sequence. It has a FIFO (First In First Out) order.
*/
package queue

import "fmt"

// Queue structure
type Queue struct {
	Queue []int
	Front int
	Rear  int
	Size  int
}

// New returns a new static Queue with given capacity
func New(capacity int) *Queue {
	return &Queue{
		Front: 0,
		Rear:  -1,
		Size:  capacity - 1,
		Queue: make([]int, capacity),
	}
}

// IsFull checks if the Queue is full
func (q *Queue) IsFull() bool {
	return q.Rear == q.Size
}

// IsEmpty checks if que Queue is empty
func (q *Queue) IsEmpty() bool {
	return q.Rear == -1
}

// Enqueue adds a value to the Queue
func (q *Queue) Enqueue(n int) {
	if q.IsFull() {
		fmt.Printf("queue is full\n")
		return
	}

	q.Rear++
	q.Queue[q.Rear] = n
}
