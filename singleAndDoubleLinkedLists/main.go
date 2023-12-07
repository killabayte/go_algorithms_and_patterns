package main

type List struct {
	head *Node
	tail *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Push(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	}
}

func (n *Node) Next() *Node {
	return n.next
}

type Node struct {
	value int
	next  *Node
}

func main() {
	l := &List{}
	l.Push(1)
	l.Push(2)
	l.Push(3)
}
