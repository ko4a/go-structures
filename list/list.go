package list

import (
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
}

func (n *Node) GetData() interface{} {
	return n.data
}

type List struct {
	head   *Node
	tail   *Node
	length int
}

func New() *List {
	return &List{nil, nil, 0}
}

func (l *List) GetLength() int {
	return l.length
}

func (l *List) Append(data interface{}) {
	node := &Node{data, nil}

	if l.head == nil || l.tail == nil {
		l.insertEmpty(node)
		return
	}

	l.tail.next = node
	l.tail = node
	l.length++

}

func (l *List) Prepend(data interface{}) {
	node := &Node{data, nil}

	if l.head == nil || l.tail == nil {
		l.insertEmpty(node)
		return
	}

	node.next = l.head
	l.head = node
	l.length++
}

func (l *List) GetHead() *Node {
	return l.head
}

func (l *List) GetTail() *Node {
	return l.tail
}

func (l *List) Print() {
	head := l.head

	for head != nil {
		fmt.Printf("%v\n", head.data)
		head = head.next
	}

}

func (l *List) Reverse() {
	if l.length <= 1 {
		return
	}

	curr := l.head.next
	prev := l.head

	for curr != nil {
		l.head = curr
		curr = curr.next
		l.head.next = prev
		prev = prev.next
	}



	fmt.Println()
}

func (l *List) insertEmpty(n *Node) {
	l.length++
	l.head = n
	l.tail = n
}
