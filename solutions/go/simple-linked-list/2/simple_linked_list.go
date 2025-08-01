package linkedlist

import "errors"

type Node struct {
    value int
    next *Node
}

type List struct {
    head *Node
    size int
}

func New(elements []int) *List {
    list := List{}
    for _, element := range elements {
        list.Push(element)
    }
    return &list
}

func (l *List) Size() int {
    return l.size
}

func (l *List) Push(element int) {
    node := Node{element, l.head}
    l.head = &node    
    l.size++
}

func (l *List) Pop() (int, error) {
    if l.size == 0 {
        return 0, errors.New("The list is empty")
    }
    node := l.head
    l.head = l.head.next
    l.size--
    return node.value, nil
}

func (l *List) Array() []int {
    array := make([]int, l.size)
    index := l.size-1    
    for node := l.head; node != nil; node = node.next {
        array[index] = node.value
        index--
    }
	return array
}

func (l *List) Reverse() *List {
    var pointer *Node
    for l.head != nil {
        prev := pointer
        pointer = l.head
        l.head = l.head.next
        pointer.next = prev
    }
    l.head = pointer
    return l
}
