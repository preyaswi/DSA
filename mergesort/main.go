package main

import "fmt"

func MergeSort(input []int) []int {
	length := len(input)
	if length <= 1 {
		return input
	}
	mid := length / 2
	left := MergeSort(input[:mid])
	right := MergeSort(input[mid:])
	return merge(left, right)

}
func merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}

	}
	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)
	return merged

}
func main() {
	arr := []int{3, 9, 2, 2, 35, 7}
	fmt.Println("before merge sort", arr)
	newArray := MergeSort(arr)
	fmt.Println("after merge sort", newArray)
}
