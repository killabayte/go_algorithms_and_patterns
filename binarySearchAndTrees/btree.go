package main

type Tree struct {
	node *Node
}

func (t *Tree) insert(value int) *Tree {
	if t.node == nil {
		t.node = &Node{value: value}
	} else {
		t.node.insert(value)
	}
	return t
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) insert(value int) {
	if value <= n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.insert(value)
		}
	}
}

func main() {
	t := &Tree{}
	t.insert(10).
		insert(8).
		insert(20).
		insert(9).
		insert(0).
		insert(15).
		insert(25)
}
