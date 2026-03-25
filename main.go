package main

import (
	"DoubleLinkedList/dlist"
	"fmt"
)

func main() {
	d := dlist.New[int]()
	fmt.Println("len after init", d.Len())
	d.Append(1)
	d.Append(2)

	fmt.Println("len after append 2 elements:", d.Len())
	d.Append(3)

	ch := d.Iterator()
	for i := range ch {
		fmt.Println("el:", i)
	}

	fmt.Println("index element 1 its:", d.Find(1))

	d.Prepend(0)

	ch = d.Iterator()
	for i := range ch {
		fmt.Println("el:", i)
	}

	d.Delete(2)
	fmt.Println("len after delete 1 element", d.Len())

	ch = d.Iterator()
	for i := range ch {
		fmt.Println("el:", i)
	}

	if err := d.Insert(0, 10); err != nil {
		fmt.Println(err)
	}

	if err := d.Insert(d.Len(), 5); err != nil {
		fmt.Println(err)
	}

	ch = d.Iterator()
	for i := range ch {
		fmt.Println("el after insert:", i)
	}
}
