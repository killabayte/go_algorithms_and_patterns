package main

type List struct {
	head *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Push(value int) {
	//TBD
}

type Node struct {
	value int
	next  *Node
}
