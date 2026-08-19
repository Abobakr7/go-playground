package main

import (
	"example.com/data-structures/linkedlists"
	"fmt"
)

func main() {
	sll := &linkedlists.SinglyLinkedList[int]{}
	fmt.Println(sll)

	sll.PushBack(1)
	sll.PushBack(2)
	sll.PushBack(3)
	sll.PushBack(4)
	sll.PushBack(5)
	sll.PushBack(6)
	sll.PushBack(7)

	fmt.Println(sll)
}
