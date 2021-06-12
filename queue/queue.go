/*
Queue

- Collection of entities that are maintained in a sequence. It has a FIFO (First In First Out) order.
*/
package queue

import "fmt"

// Queue structure
type Queue struct {
	Queue []int
	front int
	rear  int
	size  int
}

// New returns a new static Queue with given capacity
func New(capacity int) *Queue {
	return &Queue{
		front: 0,
		rear:  -1,
		size:  capacity - 1,
		Queue: make([]int, capacity),
	}
}

// IsFull checks if the Queue is full
func (q *Queue) IsFull() bool {
	return q.rear == q.size
}

// IsEmpty checks if que Queue is empty
func (q *Queue) IsEmpty() bool {
	return q.rear == -1
}

// Enqueue adds a value to the Queue
func (q *Queue) Enqueue(n int) {
	if q.IsFull() {
		fmt.Printf("queue is full\n")
		return
	}

	q.rear++
	q.Queue[q.rear] = n
}

// Dequeue removes the first element of the Queue and returns it.
// Also, updates the Queue with new front
func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		fmt.Printf("queue is empty\n")
		return 0
	}

	n := q.Queue[q.front]

	q.rear--

	tempQ := q.Queue[q.front+1:]

	for i, v := range tempQ {
		q.Queue[i] = v
	}

	q.Queue[q.size] = 0

	return n
}

// CheckFront checks the first element of the Queue without removing it
func (q *Queue) CheckFront() int {
	return q.Queue[q.front]
}

// CheckRear checks the last element of the Queue
func (q *Queue) CheckRear() int {
	return q.Queue[q.rear]
}
