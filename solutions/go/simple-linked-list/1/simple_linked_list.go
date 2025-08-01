package linkedlist

import "errors"

type Node struct {
    value int
    next *Node
}

type List struct {
    head *Node
    tail *Node
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
    node := Node{}
    node.value = element
    
    if l.size > 0 {
        l.tail.next = &node
        l.tail = &node
    } else {
        l.head = &node
        l.tail = &node
    }
    
    l.size++
}

func (l *List) Pop() (int, error) {
    if l.size == 0 {
        return 0, errors.New("The list is empty")
    }
    l.size--
    node := l.head
	if l.size == 0 {
        l.head = nil
        l.tail = nil
        return node.value, nil
    }

    for node.next != l.tail {
        node = node.next
    }
	l.tail = node
    node = node.next
    l.tail.next = nil
    
    return node.value, nil
}

func (l *List) Array() []int {
    array := []int{}
    node := l.head
    for node != nil {
        array = append(array, node.value)
        node = node.next
    }
	return array
}

func (l *List) Reverse() *List {
    if l.size == 0 {
        return l
    }
    newTail := l.head
    newHead := l.head
    pointer := l.head.next
    l.head.next = nil
    for pointer != nil {
        next := pointer.next
        pointer.next = newHead
        newHead = pointer
        pointer = next
    }

    l.head = newHead
    l.tail = newTail
    
    return l
}
