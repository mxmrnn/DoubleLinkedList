package dlist

type Node[T comparable] struct {
	Value T
	Prev  *Node[T]
	Next  *Node[T]
}
