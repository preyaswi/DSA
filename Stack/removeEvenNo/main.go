package main

import "fmt"

type Stack struct {
	list []int
}

func (s *Stack) Push(value int) {
	s.list = append(s.list, value)
}
func (s *Stack) Pop() int {
	if len(s.list) == 0 {
		return 0
	}
	lastElement := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return lastElement
}
func (s *Stack) IsEmpty() bool {
	return len(s.list) == 0
}
func RemoveEvenNumbers(s *Stack) {
	if s.IsEmpty() {
		return
	}
	element := s.Pop()
	RemoveEvenNumbers(s)
	if element%2 != 0 {
		s.Push(element)
	}
}
func main() {
	Stack := &Stack{}
	Stack.Push(30)
	Stack.Push(42)
	Stack.Push(5)
	Stack.Push(43)
	Stack.Push(44)
	fmt.Println("the array:", Stack.list)
	RemoveEvenNumbers(Stack)
	if Stack.IsEmpty() {
		fmt.Println("the array is empty")
	}
	fmt.Println("array after the removal of even numbers:", Stack.list)
}
