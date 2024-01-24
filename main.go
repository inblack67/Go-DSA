package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Size int
}

func (ll *LinkedList) GetFirst() int {
	return ll.Head.Data
}

func (ll *LinkedList) Display() {
	curr := ll.Head
	for {
		fmt.Println(curr.Data)
		curr = curr.Next
		if curr == nil {
			break
		}
	}
}

func (ll *LinkedList) GetLast() int {
	curr := ll.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	return curr.Data
}

func (ll *LinkedList) InsertAt(data int, pos int) {
	if pos > ll.Size || pos < 1 {
		panic("Invalid position")
	}

	node := NewNode(data)
	curr := ll.Head
	if pos == 1 {
		node.Next = curr
		ll.Head = node
		return
	}
	count := 1
	for count < pos-1 {
		curr = curr.Next
		count++
	}
	node.Next = curr.Next
	curr.Next = node
}

func (ll *LinkedList) deleteFirst() {
	curr := ll.Head
	ll.Head = curr.Next
}

// 1,2,3
// curr, prev = 1
// prev = 1, curr = 2
func (ll *LinkedList) reverse() {
	curr := ll.Head
	var prev *Node
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	ll.Head = prev

}

func NewNode(data int) *Node {
	node := &Node{Data: data}
	return node
}

func MakeLinkedList(arr []int) *LinkedList {
	node := NewNode(arr[0])
	ll := &LinkedList{Head: node, Size: 1}
	curr := ll.Head
	for i := 1; i < len(arr); i++ {
		node = NewNode(arr[i])
		curr.Next = node
		curr = curr.Next
		ll.Size++
	}
	return ll
}

func main() {
	// node := NewNode(1)
	// ll := &LinkedList{Head: node}
	// ll.Display()

	ll := MakeLinkedList([]int{1, 2, 3, 4, 5, 6})
	// ll.InsertAt(7, 1)
	// ll.deleteFirst()
	ll.reverse()
	ll.Display()
	// fmt.Println(ll.GetLast())

}
