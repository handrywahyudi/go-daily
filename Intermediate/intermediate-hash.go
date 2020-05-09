package main

type Node struct {
	Value int
	Next  *Node
}

type HashTable struct {
	Table map[int]*Node

	Size int
}
