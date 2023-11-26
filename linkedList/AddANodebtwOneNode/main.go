package main


import "fmt"


type Node struct {
    data int
    next *Node
}


func InsertNodeAfter(head *Node, value int, newData int) *Node {
    currentNode := head


    for currentNode != nil {
        if currentNode.data == value {
            newNode := &Node{
                data: newData,
                next: currentNode.next,
            }
            currentNode.next = newNode
            break
        }
        currentNode = currentNode.next
    }


    return head
}


func InsertNodeBefore(head *Node, value int, newData int) *Node {
    if head == nil {
        return head
    }


    if head.data == value {
        newNode := &Node{
            data: newData,
            next: head,
        }
        return newNode
    }


    currentNode := head
    var prevNode *Node


    for currentNode != nil {
        if currentNode.data == value {
            newNode := &Node{
                data: newData,
                next: currentNode,
            }
            prevNode.next = newNode
            break
        }
        prevNode = currentNode
        currentNode = currentNode.next
    }


    return head
}


func PrintLinkedList(head *Node) {
    currentNode := head
    for currentNode != nil {
        fmt.Printf("%d ", currentNode.data)
        currentNode = currentNode.next
    }
    fmt.Println()
}


func main() {
    var head *Node


    // Create a linked list: 10 -> 20 -> 30 -> 40 -> 50
    head = &Node{data: 10}
    head.next = &Node{data: 20}
    head.next.next = &Node{data: 30}
    head.next.next.next = &Node{data: 40}
    head.next.next.next.next = &Node{data: 50}


    fmt.Println("Linked List (before insertion):")
    PrintLinkedList(head)


    head = InsertNodeAfter(head, 20, 25)


    fmt.Println("Linked List (after insertion after node with value 20):")
    PrintLinkedList(head)


    head = InsertNodeBefore(head, 10, 5)


    fmt.Println("Linked List (after insertion before node with value 10):")
    PrintLinkedList(head)
}


 