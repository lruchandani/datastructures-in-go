package queue

import (
	"errors"
	"fmt"
)

//Node struct
type Node struct {
	element int
	next    *Node
}

// MyQueue type
type MyQueue struct {
	size int
	head *Node
	tail *Node
}

// Add element to the Queue
func (queue *MyQueue) Add(element int) {
	node := &Node{element, nil}
	if queue.size == 0 {
		queue.head = node
	} else {
		(*queue.tail).next = node
	}
	queue.tail = node
	queue.size++
}

// Size of queue
func (queue *MyQueue) Size() int {
	return queue.size
}

// Print queue
func (queue *MyQueue) Print() {
	var h *Node = queue.head
	for h != nil {
		fmt.Print((*h).element, " ")
		h = (*h).next
	}
	fmt.Println()
}

//Remove Element
func (queue *MyQueue) Remove() (int, error) {
	if queue.size == 0 {
		e := errors.New("Queue is empty")
		return 0, e
	}

	n := (*queue.head).element
	queue.head = (*queue.head).next
	queue.size--
	return n, nil
}
