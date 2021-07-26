package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Stack interface {
	isEmpty() bool
	push(node Node)
	pop() (Node, bool)
	peek() (Node, bool)
}

type MyStack struct {
	top *Node
}

func (myStack *MyStack) isEmpty() bool {
	return myStack.top == nil
}

func (myStack *MyStack) pop() (Node, bool) {
	if myStack.top == nil {
		return Node{}, false
	}
	poppedElement := *myStack.top
	myStack.top = myStack.top.next
	return poppedElement, true
}

func (myStack *MyStack) peek() (Node, bool) {
	if myStack.top == nil {
		return Node{}, false
	}
	return *myStack.top, true
}

func (myStack *MyStack) push(node Node) {
	node.next = myStack.top
	myStack.top = &node
}

func main() {
	myStack := MyStack{nil}
	myStack.push(Node{5, myStack.top})
	myStack.push(Node{7, myStack.top})
	myStack.push(Node{9, myStack.top})
	myStack.push(Node{11, myStack.top})
	myStack.push(Node{13, myStack.top})
	myStack.push(Node{15, myStack.top})
	myStack.push(Node{17, myStack.top})


	for !myStack.isEmpty() {
		poppedElement, _ := myStack.pop()
		fmt.Println(poppedElement.data)
	}

	fmt.Println();

	myStack.push(Node{50, myStack.top})
	myStack.push(Node{70, myStack.top})
	myStack.push(Node{90, myStack.top})
	myStack.push(Node{110, myStack.top})
	myStack.push(Node{130, myStack.top})
	myStack.push(Node{150, myStack.top})
	myStack.push(Node{170, myStack.top})


	for !myStack.isEmpty() {
		poppedElement, _ := myStack.pop()
		fmt.Println(poppedElement.data)
	}
}
