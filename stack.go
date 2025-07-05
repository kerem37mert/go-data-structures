package main

type Node struct {
	value int32
	next  *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) push(value int32) {
	newNode := &Node{value, s.top}
	s.top = newNode
}

func (s *Stack) pop() *Node {
	if s.top == nil {
		return nil
	}
	node := s.top
	s.top = s.top.next
	return node
}

func (s *Stack) peek() *Node {
	return s.top
}

func newStack() *Stack {
	return &Stack{}
}
