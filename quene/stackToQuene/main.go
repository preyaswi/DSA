package main

import "fmt"

func main() {
	stack := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("the stack", stack)
	quene := StackToQuene(stack)
	fmt.Println("the quene", quene)

}
func StackToQuene(stack []int) map[int]int {
	quene := make(map[int]int)
	for i, v := range stack {
		quene[len(stack)-i-1] = v
	}
	return quene
}
