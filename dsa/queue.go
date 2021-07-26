package main

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

type Queue interface {
	offer(node Node)
	poll() (Node, bool)
	peek() (Node, bool)
	isEmpty() bool
}

type MyQueue struct {
	front *Node
	rear *Node
}

func (myQueue *MyQueue) offer(node Node) {
	node.prev = myQueue.rear
	node.next = nil

	if myQueue.rear == nil {
		myQueue.front = &node
	} else {
			myQueue.rear.next = &node
	}
		myQueue.rear = &node
}

func (myQueue *MyQueue) poll() (Node, bool) {
	if myQueue.isEmpty() {
		return Node{}, false
	}

	polledNode := myQueue.front
	myQueue.front = myQueue.front.next

	if myQueue.isEmpty() {
			myQueue.rear = nil
	} else {
		myQueue.front.prev = nil
	}
	return *polledNode, true
}

func (myQueue *MyQueue) peek() (Node, bool) {
	if myQueue.isEmpty() {
		return Node{}, false
	}
	return *myQueue.front, true
}

func (myQueue *MyQueue) isEmpty() bool {
	return myQueue.front == nil
}

func main() {
	myQueue := MyQueue{nil, nil}
	myQueue.offer(Node{1, nil, nil})
	myQueue.offer(Node{2, nil, nil})
	myQueue.offer(Node{3, nil, nil})
	myQueue.offer(Node{4, nil, nil})
	myQueue.offer(Node{5, nil, nil})
	myQueue.offer(Node{6, nil, nil})
	
	for !myQueue.isEmpty() {
		polledElement,_ := myQueue.poll()
		fmt.Println(polledElement.data)
	}
	fmt.Println()

	myQueue.offer(Node{10, nil, nil})
	myQueue.offer(Node{20, nil, nil})
	myQueue.offer(Node{30, nil, nil})
	myQueue.offer(Node{40, nil, nil})
	myQueue.offer(Node{50, nil, nil})
	myQueue.offer(Node{60, nil, nil})
	
	for !myQueue.isEmpty() {
		polledElement,_ := myQueue.poll()
		fmt.Println(polledElement.data)
	}
}
