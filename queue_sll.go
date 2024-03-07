package main

import (
	"fmt"
)

type Node struct {
	Data interface{}
	Next *Node
}

// FIFO (enqueue,dequeue, isempty,front,rear)

// LinkedListQueue represents a queue implemented using a singly-linked list.
type LinkedListQueue struct {
	front *Node
	rear  *Node
	size  int
}

func NewLinkedListQueue() *LinkedListQueue {
	q := &LinkedListQueue{
		front: nil,
		rear:  nil,
		size:  0,
	}
	return q
}

// link list     [1]->[2]->[3]->[4]->[5]-> [6](upcoming node)
//               front                rear   node

func (q *LinkedListQueue) Enqueue(data interface{}) {

	newnode := &Node{Data: data}
	// [6] -- node
	//rear=nil	// rear
	//front=nil	// front
	// if this is the first node
	if q.rear == nil && q.front == nil {
		q.rear = newnode
		q.front = newnode
	} //rear             //rear
	// [1]->[2]->[3]->[4]->[5]->next = node [6](upcoming node)
	//                      q

	// [1]->[2]->[3]->[4]->[5]->[6]
	// front                    rear
	q.rear.Next = newnode
	q.rear = newnode
	q.size++
}

func (q *LinkedListQueue) Dequeue() (data interface{}, err error) {
	if q.size == 0 {
		return nil, fmt.Errorf("queue is empty")
	}

	// [1]->[2]->[3]->[4]->[5]->[6]
	// front                    rear
	data = q.front.Data //---> data = 1
	q.front = q.front.Next
	// [1]->[2]->[3]->[4]->[5]->[6]
	//      front               rear
	q.size--
	if q.front == nil {
		q.rear = nil
	}
	return data, nil
}

func (q *LinkedListQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *LinkedListQueue) Front() (data interface{}, err error) {
	if q.size == 0 {
		return nil, fmt.Errorf("queue is empty")
	}
	data = q.front.Data
	return data, nil
}

func (q *LinkedListQueue) Rear() (data interface{}, err error) {
	if q.size == 0 {
		return nil, fmt.Errorf("queue is empty")
	}
	data = q.rear.Data
	return data, nil
}

func main() {
	fmt.Println("Implementing queue using singly-linklist")
	queue := NewLinkedListQueue()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	f1, _ := queue.Front()
	r1, _ := queue.Rear()
	fmt.Println("Front:", f1) // Should print 1
	fmt.Println("Rear:", r1)  // Should print 3
	dq, _ := queue.Dequeue()
	fmt.Println("Dequeue:", dq) // Should print 1
	df, _ := queue.Front()
	fmt.Println("Front after dequeue:", df) // Should print 2

	fmt.Println("Is empty?", queue.IsEmpty()) // Should print false
}
