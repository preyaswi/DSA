package main

import "fmt"

type Quene struct {
	data []int
}

func (q *Quene) EnQuene(element int) {
	q.data = append(q.data, element)
}
func (q *Quene) DeQuene() int {
	if len(q.data) == 0 {
		return 0
	}
	element := q.data[0]
	q.data = q.data[1:]
	return element
}
func (q *Quene) Display() {
	fmt.Println("Quene elements:", q.data)
}
func main() {
	quene := Quene{}
	quene.EnQuene(30)
	quene.EnQuene(23)
	quene.EnQuene(43)
	quene.EnQuene(54)
	quene.EnQuene(3)
	quene.Display()
	quene.DeQuene()
	quene.Display()

}
