package binarytree

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func (bt *BinaryTree) insert(data int) {
	node := &Node{data: data}
	if bt.root == nil {
		bt.root = node
	} else {
		insertNode(bt.root, node)
	}
}

func insertNode(root *Node, node *Node) {
	if node.data <= root.data {
		if root.left == nil {
			root.left = node
		} else {
			insertNode(root.left, node)
		}
	} else {
		if root.right == nil {
			root.right = node
		} else {
			insertNode(root.right, node)
		}
	}
}

func preOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.data)
	preOrder(node.left)
	preOrder(node.right)
}

func inOrder(node *Node) {
	if node != nil {
		inOrder(node.left)
		fmt.Println(node.data)
		inOrder(node.right)
	}
}
