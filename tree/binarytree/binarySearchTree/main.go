package main

import (
	"fmt"
	"math"
)

type Node struct {
	left  *Node
	val   int
	right *Node
}
type BinarySearchTree struct {
	root *Node
}

func main() {
	bst := &BinarySearchTree{}
	bst.root = AddNode(bst.root, 11)
	bst.root = AddNode(bst.root, 10)
	bst.root = AddNode(bst.root, 51)
	bst.root = AddNode(bst.root, 71)
	bst.root = AddNode(bst.root, 111)
	bst.root = AddNode(bst.root, 90)
	fmt.Print("preorder:")
	bst.preorder(bst.root)
	fmt.Println()
	fmt.Print("inorder:")
	bst.inoder(bst.root)
	fmt.Println()
	fmt.Print("postorder:")
	bst.postorder(bst.root)
	fmt.Println()
	bst.root = delete(bst.root, 11)
	bst.contains(11)
	bst.preorder(bst.root)
	fmt.Println()
	target := 11.23
	closestval := Closest(bst.root, target)
	fmt.Printf("the closest of %f is %d", target, closestval)
	fmt.Println()
	fmt.Println("is bst:", IsBst(bst.root, 10, 111))
	fmt.Println("is 6 is there:",search(bst.root,6))

}
func AddNode(root *Node, val int) *Node {
	if root == nil {
		root = &Node{val: val, left: nil, right: nil}
		return root
	}
	if val < root.val {
		root.left = AddNode(root.left, val)
	} else {
		root.right = AddNode(root.right, val)
	}
	return root
}
func (bst *BinarySearchTree) contains(val int) {
	temp := bst.root
	if temp == nil {
		fmt.Println("tree is empty")
		return
	}
	for temp != nil {

		if val < temp.val {
			temp = temp.left
		} else if val > temp.val {
			temp = temp.right
		} else {
			fmt.Print("match found")
			return
		}
	}
	fmt.Println("match not found")

}
func delete(node *Node, val int) *Node {
	if node == nil {
		return node
	}
	if val < node.val {
		node.left = delete(node.left, val)
	} else if val > node.val {
		node.right = delete(node.right, val)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		} else {
			successorval := successor(node)
			node.val = successorval.val
			node.right = delete(node.right, successorval.val)

		}
	}
	return node

}
func successor(node *Node) *Node {
	c := node.right
	for c != nil && c.left != nil {
		c = c.left
	}
	return c
}
func (bst *BinarySearchTree) preorder(node *Node) {
	if node != nil {
		fmt.Print(node.val, " ")
		bst.preorder(node.left)
		bst.preorder(node.right)

	}
}
func (bst *BinarySearchTree) inoder(node *Node) {
	if node != nil {
		bst.inoder(node.left)
		fmt.Print(node.val, " ")
		bst.inoder(node.right)
	}
}
func (bst *BinarySearchTree) postorder(node *Node) {
	if node != nil {
		bst.postorder(node.left)
		bst.postorder(node.right)
		fmt.Print(node.val, " ")
	}
}
func Closest(node *Node, data float64) int {
	close := node.val
	mindiffe := math.Abs(float64(close) - data)
	for node != nil {
		if math.Abs(float64(node.val)-data) < mindiffe {
			mindiffe = math.Abs(float64(node.val) - data)
			close = node.val
		}
		if data < float64(node.val) {
			node = node.left
		} else {
			node = node.right
		}
	}
	return close

}
func IsBst(node *Node, min, max int) bool {
	if node == nil {
		return true
	}
	return node.val >= min && node.val <= max &&
		IsBst(node.left, min, node.val) &&
		IsBst(node.right, node.val, max)
}
func search(root *Node, data int) bool {
    if root == nil {
        return false
    }
    if root.val == data {
        return true
    }
    if root.val > data {
        return search(root.left, data)
    } else {
        return search(root.right, data)
    }
}
