package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func addNode(t *Node, v int) int {
	if root == nil {
		t = &Node{v, nil}
		root = t
		return 0
	}

	if v == t.Value {
		fmt.Println("Node already exist:", v)
		return -1
	}

	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}
	return addNode(t.Next, v)
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty List.")
		return
	}

	for t != nil {
		fmt.Print("%d-> ", t.Value)
		t = t.Next
	}

	fmt.Println()
}

var root = new(Node)

func main() {
	fmt.Println(root)
	root = nil
	fmt.Println(root)
}
