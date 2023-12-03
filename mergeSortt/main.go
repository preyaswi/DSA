// checking how many inversions have been done(a[i]>a[j],i<j)
package main

import "fmt"

func main() {
	arr := []int{2, 4, 1, 3, 5}
	fmt.Println(arr)
	inversions, merged := MergeSorting(arr)
	fmt.Println(merged)
	fmt.Println(inversions)

}
func MergeSorting(arr []int) (int, []int) {
	length := len(arr)
	if length <= 1 {
		return 0, arr
	}
	mid := length / 2
	leftinversion, left := MergeSorting(arr[:mid])
	rightInversion, right := MergeSorting(arr[mid:])
	mergerdInversion, mergedSort := mergeAndCount(left, right)
	totalinversion := leftinversion + rightInversion + mergerdInversion
	return totalinversion, mergedSort
}
func mergeAndCount(left, right []int) (int, []int) {
	var inversions int
	var merged []int

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			inversions += len(left) - i
			merged = append(merged, right[j])
			j++
		}
	}

	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	return inversions, merged
}
