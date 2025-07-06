package main

type Queue struct {
	front *Node
	rear  *Node
}

func (q *Queue) enqueue(value int32) {
	newNode := &Node{value, nil}

	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
		return
	}

	q.rear.next = newNode
	q.rear = newNode
}

func (q *Queue) dequeue() {
	if q.front == nil {
		return
	}

	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	}
}

func newQueue() *Queue {
	return &Queue{}
}
