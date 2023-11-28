package main

import "fmt"

func main() {
	arr := []int{5, 67, 8, 3, 2, 5, 7}
	fmt.Println("array before quicksort", arr)
	newArray := quickSort(arr)
	fmt.Println("array after quicksort", newArray)
}
func quickSort(input []int) []int {
	if len(input) <= 1 {
		return input
	}
	var left, right []int
	pivot := input[len(input)-1]
	for i := 0; i < len(input)-1; i++ {
		if input[i] <= pivot {
			left = append(left, input[i])
		} else {
			right = append(right, input[i])
		}

	}
	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, pivot), right...)
}
