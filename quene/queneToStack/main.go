package main

import "fmt"

func main() {
	quene := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("the quene:", quene)
	stack:=QueneToStack(quene)
	fmt.Println("the stack",stack)
}
func QueneToStack(quene []int) map[int]int {
	stack := make(map[int]int)
	for i, v := range quene {
		stack[i] = v
	}
	return stack
}
