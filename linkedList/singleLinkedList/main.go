package main


import (
    "fmt"
)


// single linked list
type Node struct {
    data int
    next *Node
}
type Linkedalist struct {
    head *Node
}


func (link *Linkedalist) Insert(data int) {
    newNode := &Node{data: data}
    if link.head == nil {
        link.head = newNode
    } else {
        current := link.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode


    }
}
func (link Linkedalist) Display() {
    current := link.head
    if current == nil {
        fmt.Println("linked list is empty")
        return
    }
    fmt.Println("Linked list is ")
    for current != nil {
        fmt.Printf("%d ", current.data)
        current = current.next
    }
    fmt.Println()
}
func main() {
    list := Linkedalist{}
    list.Insert(32)
    list.Insert(43)
    list.Insert(45)
    list.Display()
}
