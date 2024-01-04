package main

import (
	"fmt"

	
)

type node struct {
	data     int
	nextNode *node
}
type queue struct {
	head *node
	tail *node
}

func main() {
var que queue=queue{}
que.EnQueue(14)
que.EnQueue(142)
que.EnQueue(141)
que.EnQueue(1)
printQueue(&que)
que.Dequeue()
printQueue(&que)

}
func (q *queue) EnQueue(element int) {
	newNode := &node{data: element, nextNode: nil}
	if q.head == nil {
		q.head = newNode
		q.tail=newNode

	} else {
		q.tail.nextNode = newNode

		q.tail = newNode
	}
}
func (q *queue)Dequeue()  {
	if q.head==nil{
		return
	}
	q.head=q.head.nextNode
}
func printQueue(q *queue)  {
	if q.head==nil{
		fmt.Println("the queue is empty")
		return
	}
	currentNode:=q.head
	for currentNode!=nil{
		fmt.Print(currentNode.data," ")
		currentNode=currentNode.nextNode
	}
	fmt.Println()
}
