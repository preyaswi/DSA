package main

import (
	"fmt"
)

type stack struct {
	list []int
}

func main() {
	stack := stack{}
	for i := 0; i < 5; i++ {
		stack.Push(i)
	}
	fmt.Println(stack.list)
	Removemiddle(&stack)
	fmt.Println(stack.list)
}
func (s *stack) Push(element int) {
	s.list = append(s.list, element)
}
func (s *stack) Pop() int {
	element := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return element
}
func (s *stack) Size() int {
	return len(s.list) / 2
}
func (s *stack) Isempty() bool {
	return len(s.list) == 0
}
func Removemiddle(stacks *stack) {
	middleInd := stacks.Size()
	tempStack := stack{}
	for i := 0; i < middleInd; i++ {
		popped := stacks.Pop()
		tempStack.Push(popped)
	}
	stacks.Pop()
	for !tempStack.Isempty() {
		popped := tempStack.Pop()
		stacks.Push(popped)
	}

}
