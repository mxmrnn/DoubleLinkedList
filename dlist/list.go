package dlist

import (
	"errors"
)

type DoubleLinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	len  int
}

var ErrIndexOutOfRange = errors.New("index out of range")

func (d *DoubleLinkedList[T]) initFirst(value T) {
	n := &Node[T]{Value: value}
	d.head = n
	d.tail = n
	d.len = 1
}

func (d *DoubleLinkedList[T]) Append(value T) {

	if d.head == nil && d.tail == nil {
		d.initFirst(value)
		return
	}

	n := &Node[T]{
		Value: value,
		Prev:  d.tail,
	}

	d.tail.Next = n
	d.tail = n
	d.len++
}

func (d *DoubleLinkedList[T]) Prepend(value T) {
	if d.head == nil && d.tail == nil {
		d.initFirst(value)
		return
	}

	n := &Node[T]{
		Value: value,
		Prev:  nil,
		Next:  d.head,
	}

	d.head.Prev = n
	d.head = n
	d.len++
}

func (d *DoubleLinkedList[T]) Insert(index int, value T) error {
	if index < 0 || index > d.len {
		return ErrIndexOutOfRange
	}

	if index == 0 {
		d.Prepend(value)
		return nil
	}

	if index == d.Len() {
		d.Append(value)
		return nil
	}

	cur := d.head

	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	n := &Node[T]{
		Value: value,
		Prev:  cur.Prev,
		Next:  cur,
	}

	cur.Prev.Next = n
	cur.Prev = n

	d.len++
	return nil
}

func (d *DoubleLinkedList[T]) Delete(value T) {
	cur := d.head

	for cur != nil {
		if cur.Value == value {
			break
		}
		cur = cur.Next
	}

	if cur == nil {
		return
	}

	if d.Len() == 1 {
		d.head = nil
		d.tail = nil
		d.len--
		return
	}

	if cur == d.head {
		d.head = cur.Next
		d.head.Prev = nil
	}

	if cur == d.tail {
		d.tail = cur.Prev
		d.tail.Next = nil
	}

	if cur.Prev != nil {
		cur.Prev.Next = cur.Next
	}

	if cur.Next != nil {
		cur.Next.Prev = cur.Prev
	}

	d.len--
}

func (d *DoubleLinkedList[T]) Find(value T) int {
	cur := d.head

	i := 0
	for cur != nil {
		if cur.Value == value {
			return i
		}
		cur = cur.Next
		i++
	}
	return -1
}

func (d *DoubleLinkedList[T]) Len() int {
	return d.len
}

func (d *DoubleLinkedList[T]) Iterator() (ch chan any) {
	ch = make(chan any)

	go func(d *DoubleLinkedList[T]) {
		cur := d.head

		for cur != nil {
			ch <- cur.Value
			cur = cur.Next
		}
		close(ch)
	}(d)

	return ch
}

func New[T comparable]() (dl DoubleLinkedList[T]) {
	return DoubleLinkedList[T]{}
}
