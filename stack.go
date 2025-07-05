package main

type Node struct {
	value int32
	next  *Node
}

type Stack struct {
	head *Node
}

func (s *Stack) push(value int32) {
	newNode := &Node{value, s.head}
	s.head = newNode
}

func (s *Stack) pop() *Node {
	if s.head == nil {
		return nil
	}
	node := s.head
	s.head = s.head.next
	return node
}

func (s *Stack) peek() *Node {
	return s.head
}

func newStack() *Stack {
	return &Stack{}
}
