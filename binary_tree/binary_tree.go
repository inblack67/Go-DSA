package binarytree

import "fmt"

//     5
//    / \
//   3   7
//  / \
// 2   4

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	root           *Node
	nodeToRootPath NodeToRootPath
}

type NodeToRootPath []*Node

func (bt *BinaryTree) Insert(data int) {
	node := &Node{Data: data}
	if bt.root == nil {
		bt.root = node
	} else {
		InsertNode(bt.root, node)
	}
}

func InsertNode(root *Node, node *Node) {
	if node.Data <= root.Data {
		if root.Left == nil {
			root.Left = node
		} else {
			InsertNode(root.Left, node)
		}
	} else {
		if root.Right == nil {
			root.Right = node
		} else {
			InsertNode(root.Right, node)
		}
	}
}

func PreOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.Data)
	PreOrder(node.Left)
	PreOrder(node.Right)
}

func InOrder(node *Node) {
	if node != nil {
		InOrder(node.Left)
		fmt.Println(node.Data)
		InOrder(node.Right)
	}
}

func CalculateSize(root *Node) int {
	if root == nil {
		return 0
	}
	lhs := CalculateSize(root.Left)
	rhs := CalculateSize(root.Right)
	return lhs + rhs + 1
}

func CalculateSum(root *Node) int {
	if root == nil {
		return 0
	}
	lhs := CalculateSum(root.Left)
	rhs := CalculateSum(root.Right)
	return lhs + rhs + root.Data
}

func CalculateMax(root *Node) int {
	if root == nil {
		return 0
	}
	max := root.Data
	lhs := CalculateMax(root.Left)
	rhs := CalculateMax(root.Right)
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

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func CalculateHeight(root *Node) int {
	if root == nil {
		return 0
	}
	lhs := CalculateHeight(root.Left)
	rhs := CalculateHeight(root.Right)
	return Max(lhs, rhs) + 1
}

func (bt *BinaryTree) CalculateNodeToRootPath(root *Node, data int) bool {
	if root == nil {
		return false
	}
	if root.Data == data {
		bt.nodeToRootPath = append(bt.nodeToRootPath, root)
		return true
	}
	lhs := bt.CalculateNodeToRootPath(root.Left, data)
	rhs := bt.CalculateNodeToRootPath(root.Right, data)
	if lhs {
		bt.nodeToRootPath = append(bt.nodeToRootPath, root)
		return true
	} else if rhs {
		bt.nodeToRootPath = append(bt.nodeToRootPath, root)
		return true
	}
	return false
}

// func main() {
// 	bt := &BinaryTree{}
// 	bt.Insert(5)
// 	bt.Insert(3)
// 	bt.Insert(7)
// 	bt.Insert(2)
// 	bt.Insert(4)
// 	// preOrder(bt.root)
// 	// inOrder(bt.root)
// 	// fmt.Println(calculateHeight(bt.root))
// 	bt.CalculateNodeToRootPath(bt.root, 4)
// 	for _, v := range bt.nodeToRootPath {
// 		fmt.Println(v.Data)
// 	}
// }
