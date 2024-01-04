package main

import "fmt"

type Node struct {
	Left  *Node
	Val   int
	Right *Node
}
type BinaryTree struct {
	Root *Node
}

func main() {
	node1 := &Node{nil, 1, nil}
	node2 := &Node{nil, 2, nil}
	node3 := &Node{nil, 3, nil}
	node4 := &Node{nil, 4, nil}
	node5 := &Node{nil, 5, nil}
	tree := &BinaryTree{}
	tree.Root = node1
	tree.Root.Left = node2
	tree.Root.Right = node3
	tree.Root.Left.Left = node4
	tree.Root.Left.Right = node5
	fmt.Println("In-Order traversal:")
	inorder(tree.Root)
	fmt.Println("Pre-Order traversal:")
	preorder(tree.Root)
	fmt.Println("Post-Order traversal:")
	Postorder(tree.Root)
}

func inorder(root *Node) {
	if root != nil {
		inorder(root.Left)
		fmt.Println(root.Val)
		inorder(root.Right)
	}
}
func preorder(root *Node) {
	if root != nil {
		fmt.Println(root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
}
func Postorder(root *Node) {
	if root != nil {
		Postorder(root.Left)
		Postorder(root.Right)
		fmt.Println(root.Val)
	}
}
