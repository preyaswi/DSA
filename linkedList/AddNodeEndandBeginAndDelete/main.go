package main


import "fmt"


type Node struct {
    data int
    next *Node
}


func AddnodeAtBegining(head *Node, data int) *Node {
    newNode := &Node{data: data,
        next: head}
    return newNode
}
func AddNodeAtLast(head *Node, data int) *Node {
    newNode := &Node{data: data,
        next: nil}
    if head == nil {
        return newNode
    } else {
        currentNode := head
        for currentNode.next != nil {
            currentNode = currentNode.next
        }
        currentNode.next = newNode
    }
    return head
}
func PrintLinkedList(head *Node) {
    current := head
    if current == nil {
        fmt.Println("no linked list")
        return
    } else {
        for current != nil {
            fmt.Printf("%d ", current.data)
            current = current.next
            fmt.Println()
        }
    }
}
func DeleteANode(head *Node, value int) *Node {
    if head == nil {
        fmt.Println("no value to be deleted")
    }
    if head != nil && head.data == value {
        return head.next
    }
    current := head
    var prevNode *Node
    for current != nil && current.data != value {
        prevNode = current
        current = current.next
    }
    if current != nil {
        prevNode.next = current.next
    }
    return head
}
func main() {
    var head *Node
    head = AddNodeAtLast(head, 12)
    head = AddNodeAtLast(head, 34)
    head = AddNodeAtLast(head, 45)
    fmt.Println("node added at the last of the linked list")
    PrintLinkedList(head)
    head = AddnodeAtBegining(head, 1)
    fmt.Println("node added at the begining of the linked list")
    PrintLinkedList(head)
    head = DeleteANode(head, 34)
    fmt.Println(" linked list after deleting")
    PrintLinkedList(head)
}
