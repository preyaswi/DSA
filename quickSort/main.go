package main

import "fmt"

//first method
// func main() {
// 	arr := []int{5, 67, 8, 3, 2, 5, 7}
// 	fmt.Println("array before quicksort", arr)
// 	newArray := quickSort(arr)
// 	fmt.Println("array after quicksort", newArray)
// }
// func quickSort(input []int) []int {
// 	if len(input) <= 1 {
// 		return input
// 	}
// 	var left, right []int
// 	pivot := input[len(input)-1]
// 	for i := 0; i < len(input)-1; i++ {
// 		if input[i] <= pivot {
// 			left = append(left, input[i])
// 		} else {
// 			right = append(right, input[i])
// 		}

// 	}
// 	left = quickSort(left)
// 	right = quickSort(right)
// 	return append(append(left, pivot), right...)
// }
func main() {
	arr := []int{2, 4, 56, 3, 2, 1, 5, 6, 4, 2, 5}
	quickSort(arr, 0, 10)

}
func quickSort(arr []int, startInd, lastInd int) {
	fmt.Println("array before sorting:", arr)
	quickSortMethod(arr, startInd, lastInd)
	fmt.Println("array after sorting:", arr)

}
func quickSortMethod(arr []int, startInd, lastInd int) {
	if lastInd <= startInd {
		return
	}
	pivotInd := startInd
	leftInd := startInd + 1
	rightInd := lastInd
	for rightInd >= leftInd {
		if arr[rightInd] < arr[pivotInd] && arr[leftInd] > arr[pivotInd] {
			swap(arr, rightInd, leftInd)
			leftInd++
			rightInd--
		}
		if arr[leftInd] <= arr[pivotInd] {
			leftInd++
		}
		if arr[rightInd] >= arr[pivotInd] {
			rightInd--
		}
	}
	swap(arr, rightInd, pivotInd)
	quickSortMethod(arr, startInd, rightInd-1)
	quickSortMethod(arr, rightInd+1, lastInd)

}
func swap(arr []int, a int, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}
