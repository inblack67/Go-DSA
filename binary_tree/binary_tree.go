package binarytree

import "fmt"

//     5
//    / \
//   3   7
//  / \
// 2   4

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root           *Node
	nodeToRootPath NodeToRootPath
}

type NodeToRootPath []*Node

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

func calculateSize(root *Node) int {
	if root == nil {
		return 0
	}
	lhs := calculateSize(root.left)
	rhs := calculateSize(root.right)
	return lhs + rhs + 1
}

func calculateSum(root *Node) int {
	if root == nil {
		return 0
	}
	lhs := calculateSum(root.left)
	rhs := calculateSum(root.right)
	return lhs + rhs + root.data
}

func calculateMax(root *Node) int {
	if root == nil {
		return 0
	}
	max := root.data
	lhs := calculateMax(root.left)
	rhs := calculateMax(root.right)
	var candidate int
	if lhs > max {
		candidate = lhs
	} else {
		candidate = rhs
	}
	if max > candidate {
		return max
	} else {
		return candidate
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calculateHeight(root *Node) int {
	if root == nil {
		return 0
	}
	lhs := calculateHeight(root.left)
	rhs := calculateHeight(root.right)
	return max(lhs, rhs) + 1
}

func (bt *BinaryTree) calculateNodeToRootPath(root *Node, data int) bool {
	if root == nil {
		return false
	}
	if root.data == data {
		bt.nodeToRootPath = append(bt.nodeToRootPath, root)
		return true
	}
	lhs := bt.calculateNodeToRootPath(root.left, data)
	rhs := bt.calculateNodeToRootPath(root.right, data)
	if lhs {
		bt.nodeToRootPath = append(bt.nodeToRootPath, root)
		return true
	} else if rhs {
		bt.nodeToRootPath = append(bt.nodeToRootPath, root)
		return true
	}
	return false
}

func main() {
	bt := &BinaryTree{}
	bt.insert(5)
	bt.insert(3)
	bt.insert(7)
	bt.insert(2)
	bt.insert(4)
	// preOrder(bt.root)
	// inOrder(bt.root)
	// fmt.Println(calculateHeight(bt.root))
	bt.calculateNodeToRootPath(bt.root, 4)
	for _, v := range bt.nodeToRootPath {
		fmt.Println(v.data)
	}
}
