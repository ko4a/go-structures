package list

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type List struct {
	head   *Node
	tail   *Node
	length int
}

func New() *List {
	return &List{nil, nil, 0}
}

func (l *List) getLength() int  {
	return l.length
}

func (l *List) Append(data interface{}) {
	node := &Node{data, nil}

	l.length++

	if l.head == nil || l.tail == nil {
		l.head = node
		l.tail = node
		return
	}

	l.tail.next = node
	l.tail = node

}

func (l *List) Print() {
	head := l.head

	for head != nil {
		fmt.Printf("%v\n", head.data)
		head = head.next
	}

}
