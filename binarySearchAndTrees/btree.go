package main

type Tree struct {
	node *Node
}

type Node struct {
	value int
	left  *Node
	right *Node
}
