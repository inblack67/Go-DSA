package ll

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

func (ll *LinkedList) InsertAtLast(data int) {
	node := NewNode(data)
	if ll.Head == nil {
		ll.Head = node
		ll.Size++
	} else {
		curr := ll.Head
		for {
			if curr.Next == nil {
				break
			}
			curr = curr.Next
		}
		curr.Next = node
	}
}

func (ll *LinkedList) DeleteFirst() {
	curr := ll.Head
	ll.Head = curr.Next
}

// 1,2,3
// curr, prev = 1
// prev = 1, curr = 2
func (ll *LinkedList) Reverse() {
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

// 2, 4, 3
//    6, 4

func addHelper(node1, node2 *Node, pv1, pv2 int, ll3 *LinkedList) int {

	if node1 == nil && node2 == nil {
		return 0 // 0 carry
	}

	if pv1 > pv2 {
		carry := addHelper(node1.Next, node2, pv1-1, pv2, ll3)
		sum := node1.Data + carry
		data := sum % 10
		ll3.InsertAtLast(data)

		newCarry := sum / 10

		return newCarry
	} else if pv1 < pv2 {
		carry := addHelper(node1, node2.Next, pv1, pv2-1, ll3)
		sum := node2.Data + carry
		data := sum % 10
		ll3.InsertAtLast(data)

		newCarry := sum / 10

		return newCarry
	} else {
		// pv1 = pv2 time to sum
		carry := addHelper(node1.Next, node2.Next, pv1-1, pv2-1, ll3)
		sum := node1.Data + node2.Data + carry
		data := sum % 10
		ll3.InsertAtLast(data)
		newCarry := sum / 10

		return newCarry
	}
}

func addLinkedLists(ll1, ll2, ll3 *LinkedList) {
	carry := addHelper(ll1.Head, ll2.Head, ll1.Size, ll2.Size, ll3)
	if carry > 0 {
		ll3.InsertAt(carry, 1)
	}
}

func main() {
	ll3 := &LinkedList{}

	ll1 := MakeLinkedList([]int{2, 4, 3})
	ll2 := MakeLinkedList([]int{5, 6, 4})

	// ll1 := MakeLinkedList([]int{9, 9, 9})
	// ll2 := MakeLinkedList([]int{1})

	addLinkedLists(ll1, ll2, ll3)
	ll3.Display()
}
