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
