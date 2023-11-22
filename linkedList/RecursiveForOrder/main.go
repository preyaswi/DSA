// in recursive
package main


import (
    "fmt"
)


type Node struct {
    data int
    next *Node
}


func PrintLinkedList(head *Node) {
    fmt.Println("Linked List (Original Order):")
    printOriginalOrder(head)
    fmt.Println("\nLinked List (Reverse Order):")
    printReverseOrder(head)
    fmt.Println()
}
func printOriginalOrder(node *Node) {
    if node != nil {
        fmt.Printf("%d ", node.data)
        printOriginalOrder(node.next)
    }
}
func printReverseOrder(node *Node) {
    if node != nil {
        printReverseOrder(node.next)
        fmt.Printf("%d ", node.data)
    }
}


func main() {
    var head *Node

    head = &Node{data: 10}
    head.next = &Node{data: 20}
    head.next.next = &Node{data: 30}
    head.next.next.next = &Node{data: 40}
    head.next.next.next.next = &Node{data: 50}


    PrintLinkedList(head)
}



