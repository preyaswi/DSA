package main


import (
    "fmt"
)
type Node struct {
    data int
    next *Node
}


func ToConvertArrtoLil(arr []int) *Node {
    var head, tail *Node
    for _, val := range arr {
        newNode := &Node{data: val,
            next: nil}
        if head == nil {
            head = newNode
            tail = newNode
        } else {
            tail.next = newNode
            tail = newNode
        }
    }
    return head
}
func PrintLinkedList(head *Node) {
    current := head
    for current != nil {
        fmt.Printf("%d ", current.data)
        current = current.next
    }
    fmt.Println()
}
func main() {
    arr := []int{12, 13, 14, 15, 16}
    linkedList := ToConvertArrtoLil(arr)
    PrintLinkedList(linkedList)
}
