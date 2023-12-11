package main

import "fmt"

type Queue struct {
	data []int
}

func (q *Queue) EnQueue(element int) {
	q.data = append(q.data, element)
}
func (q *Queue) DeQueue() int {
	if len(q.data) == 0 {
		return 0
	}
	element := q.data[0]
	q.data = q.data[1:]
	return element
}
func (q *Queue) Display() {
	fmt.Println("Queue elements:", q.data)
}
func main() {
	queue := Queue{}
	queue.EnQueue(30)
	queue.EnQueue(23)
	queue.EnQueue(43)
	queue.EnQueue(54)
	queue.EnQueue(3)
	queue.Display()
	queue.DeQueue()
	queue.Display()

}
