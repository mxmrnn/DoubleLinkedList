package dlist

type DoubleLinkedList struct {
	firstNode *Node
	lastNode  *Node
	length    int
}

func (d *DoubleLinkedList) Append(value any) {

	if d.firstNode == nil && d.lastNode == nil {
		d.firstNode = &Node{
			Value: value,
			Prev:  nil,
			Next:  nil,
		}
		d.lastNode = d.firstNode
		d.length++
		return
	}

	d.lastNode.Next = &Node{
		Value: value,
		Prev:  d.lastNode,
		Next:  nil,
	}

	d.lastNode = d.lastNode.Next

	d.length++
}

func (d *DoubleLinkedList) Prepend(value any) {
	if d.firstNode == nil && d.lastNode == nil {
		d.firstNode = &Node{
			Value: value,
			Prev:  nil,
			Next:  nil,
		}
		d.lastNode = d.firstNode
		d.length++
		return
	}

	d.firstNode.Prev = &Node{
		Value: value,
		Prev:  nil,
		Next:  d.firstNode,
	}

	d.firstNode = d.firstNode.Prev

	d.length++
}

func (d *DoubleLinkedList) Insert(index int, value any) {
	if index < 0 || index > d.Len() {
		return
	}

	if index == 0 {
		d.Prepend(value)
		return
	}

	if index == d.Len() {
		d.Append(value)
		return
	}

	cur := d.firstNode

	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	now := &Node{
		Value: value,
		Prev:  cur.Prev,
		Next:  cur,
	}

	cur.Prev.Next = now
	cur.Prev = now

	d.length++
}

func (d *DoubleLinkedList) Delete(value any) {
	cur := d.firstNode

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
		d.firstNode = nil
		d.lastNode = nil
		d.length--
		return
	}

	if cur == d.firstNode {
		d.firstNode = cur.Next
	}

	if cur == d.lastNode {
		d.lastNode = cur.Prev
	}

	if cur.Prev != nil {
		cur.Prev.Next = cur.Next
	}

	if cur.Next != nil {
		cur.Next.Prev = cur.Prev
	}

	d.length--
}

func (d *DoubleLinkedList) Find(value any) int {
	cur := d.firstNode

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

func (d *DoubleLinkedList) Len() int {
	return d.length
}

func New() (dl DoubleLinkedList) {
	return DoubleLinkedList{}
}

func (d *DoubleLinkedList) Iterator() (ch chan any) {
	ch = make(chan any)

	go func(d *DoubleLinkedList) {
		cur := d.firstNode

		for cur != nil {
			ch <- cur.Value
			cur = cur.Next
		}
		close(ch)
	}(d)

	return ch
}
