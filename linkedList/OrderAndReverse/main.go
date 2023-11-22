package main


import (
    "fmt"
    "strings"
)


type Node struct {
    data int
    next *Node
}


func PrintLinkedList(head *Node) {
    currentNode := head
    var elements []string


    for currentNode != nil {
        elements = append(elements, fmt.Sprintf("%d", currentNode.data))
        currentNode = currentNode.next
    }


    fmt.Println("Linked List (Original Order):")
    fmt.Println(strings.Join(elements, " "))


    fmt.Println("Linked List (Reverse Order):")
    for i := len(elements) - 1; i >= 0; i-- {
        fmt.Printf("%s ", elements[i])
    }
    fmt.Println()
}


func main() {
    var head *Node


    // Create a linked list: 10 -> 20 -> 30 -> 40 -> 50
    head = &Node{data: 10, next: &Node{data: 20, next: &Node{data: 30, next: &Node{data: 40, next: &Node{data: 50}}}}}


    PrintLinkedList(head)
}
