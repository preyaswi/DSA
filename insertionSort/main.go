// ascending order using insertionsort
package main

import "fmt"

func main() {
	arr := []int{3, 4, 53, 2, 36, 9}
	fmt.Println("array before sorting", arr)
	insertionSort(arr)
	fmt.Println("array after sorting", arr)

}
func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}
