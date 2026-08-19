package linkedlists

import (
	"fmt"
	"strings"
)

type node[T comparable] struct {
	val  T
	next *node[T]
}

type SinglyLinkedList[T comparable] struct {
	head *node[T]
	tail *node[T]
	size int
}

func (sll *SinglyLinkedList[T]) PushFront(val T) {
	newNode := &node[T]{val: val, next: sll.head}
	sll.head = newNode

	if sll.tail == nil {
		sll.tail = sll.head
	}
	sll.size++
}

func (sll *SinglyLinkedList[T]) PushBack(val T) {
	if sll.head == nil {
		sll.head = &node[T]{val: val}
		sll.tail = sll.head
		sll.size++
		return
	}

	sll.tail.next = &node[T]{val: val}
	sll.tail = sll.tail.next
	sll.size++
}

func (sll *SinglyLinkedList[T]) PopFront() {
	if sll.head == nil {
		return
	}

	if sll.head == sll.tail {
		sll.head = nil
		sll.tail = nil
		sll.size--
		return
	}

	sll.head = sll.head.next
	sll.size--
}

func (sll *SinglyLinkedList[T]) PopBack() {
	if sll.tail == nil {
		return
	}

	if sll.head == sll.tail {
		sll.head = nil
		sll.tail = nil
		sll.size--
		return
	}

	node := sll.head
	for ; node.next != sll.tail; node = node.next {
	}
	sll.tail = node
	sll.tail.next = nil
	sll.size--
}

func (sll *SinglyLinkedList[T]) Get(idx int) (val T, ok bool) {
	if sll.head == nil {
		return val, false
	}

	if idx < 0 || idx >= sll.size {
		return val, false
	}

	node := sll.head
	for i := 0; i < idx; i++ {
		node = node.next
	}

	return node.val, true
}

func (sll *SinglyLinkedList[T]) Set(idx int, val T) bool {
	if sll.head == nil || idx < 0 || idx >= sll.size {
		return false
	}

	node := sll.head
	for i := 0; i < idx; i++ {
		node = node.next
	}
	node.val = val

	return true
}

func (sll *SinglyLinkedList[T]) Insert(idx int, val T) bool {
	if idx < 0 || idx > sll.size {
		return false
	}

	switch idx {
	case 0:
		sll.PushFront(val)
		return true
	case sll.size:
		sll.PushBack(val)
		return true
	}

	prev := sll.head
	for i := 0; i < idx-1; i++ {
		prev = prev.next
	}
	prev.next = &node[T]{val: val, next: prev.next}
	sll.size++

	return true
}

func (sll *SinglyLinkedList[T]) Remove(idx int) bool {
	if sll.head == nil || idx < 0 || idx >= sll.size {
		return false
	}

	if idx == 0 {
		sll.head = sll.head.next
		if sll.head == nil {
			sll.tail = nil
		}
		sll.size--
		return true
	}

	prev := sll.head
	for i := 0; i < idx-1; i++ {
		prev = prev.next
	}
	removed := prev.next
	prev.next = removed.next
	if removed == sll.tail {
		sll.tail = prev
	}
	removed.next = nil
	sll.size--

	return true
}

func (sll *SinglyLinkedList[T]) Contains(val T) bool {
	if sll.head == nil {
		return false
	}

	for node := sll.head; node != nil; node = node.next {
		if val == node.val {
			return true
		}
	}

	return false
}

func (sll *SinglyLinkedList[T]) IndexOf(val T) int {
	for i, node := 0, sll.head; node != nil; i, node = i+1, node.next {
		if val == node.val {
			return i
		}
	}

	return -1
}

func (sll *SinglyLinkedList[T]) Reverse() {
	var prev, cur *node[T] = nil, sll.head
	sll.tail = sll.head
	for cur != nil {
		tmp := cur.next
		cur.next = prev
		prev = cur
		cur = tmp
	}
	sll.head = prev
}

func (sll *SinglyLinkedList[T]) Len() int {
	return sll.size
}

func (sll *SinglyLinkedList[T]) String() string {
	if sll.head == nil {
		return "Empty Singly LinkedList: <nil>"
	}

	var b strings.Builder

	b.WriteString("[")
	for node := sll.head; node.next != nil; node = node.next {
		fmt.Fprintf(&b, "%v, ", node.val)
	}
	fmt.Fprintf(&b, "%v]", sll.tail.val)

	return b.String()
}
