package main

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) push(element int) {
	s.data = append(s.data, element)
}
func (s *Stack) pop() int {
	if len(s.data) == 0 {
		return 0
	}
	element := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return element
}
func (s *Stack) peek() int {
	if len(s.data) == 0 {
		return 0
	}
	return s.data[len(s.data)-1]
}
func (s *Stack) Display() {
	fmt.Println("Stack elements are:", s.data)
}
func main() {
	stack := Stack{}
	stack.push(98)
	stack.push(23)
	stack.push(32)
	stack.push(5)
	stack.push(9)
	stack.Display()
	first := stack.peek()
	fmt.Println("the first element", first)
	poppedValue := stack.pop()
	fmt.Println("the popped value", poppedValue)
	stack.Display()

}
