package main

import (
	"fmt"
	"strconv"
)

type Stack struct {
	elements []int
}

func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

func (s *Stack) Pop() int {
	if len(s.elements) == 0 {
		return -1 // Assuming -1 represents an error or an empty stack
	}
	popped := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return popped
}

func evaluatePostfix(expression string) int {
	stack := Stack{}

	for _, char := range expression {
		if char >= '0' && char <= '9' {
			digit, _ := strconv.Atoi(string(char))
			stack.Push(digit)
		} else {
			operand2 := stack.Pop()
			operand1 := stack.Pop()

			switch char {
			case '+':
				stack.Push(operand1 + operand2)
			case '-':
				stack.Push(operand1 - operand2)
			case '*':
				stack.Push(operand1 * operand2)
			case '/':
				stack.Push(operand1 / operand2)
			}
		}
	}

	return stack.Pop()
}

func main() {
	fmt.Println(evaluatePostfix("23*5+"))   // (2*3) + 5 = 11
	fmt.Println(evaluatePostfix("92+4*8+")) // ((9+2)*4)+8 =52
}
