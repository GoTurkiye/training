package main

import "fmt"

type MyItem int

type Node struct {
	data MyItem
	next *Node
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Insert(data MyItem) {
	node := &Node{
		data: data,
		next: nil,
	}
	if l.Head == nil {
		l.Head = node
		return
	}
	head := l.Head
	for head.next != nil {
		head = head.next
	}
	head.next = node
}

func (l *LinkedList) Delete(data MyItem)  {
	prev := l.Head
	curr := l.Head

	for curr != nil {
		// 1 2 3 4 5
		// 1 2   4 5
		if curr.data == data {
			prev.next = curr.next
		}
		prev = curr
		curr = curr.next
	}
}

func (l *LinkedList) Print()  {
	list := l.Head
	for list != nil {
		fmt.Printf(" -> %d", list.data)
		list = list.next
	}
	println()
}

func main() {
	l := NewLinkedList()
	l.Insert(3)
	l.Insert(5)
	l.Insert(7)
	l.Insert(7)
	l.Insert(9)
	l.Print()
	l.Delete(7)
	l.Print()
}
