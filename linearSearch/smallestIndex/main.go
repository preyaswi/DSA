package main

import "fmt"

func main() {
	array := []int{16, 5, 8, 12, 16, 23, 16}

	result := SmallestElement(array)
	if result == -1 {
		fmt.Println("element not thre")
	} else {
		fmt.Println("Smallest element is index ", result)
	}
}
func SmallestElement(arr []int) int {
	sm := arr[0]
	index := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] < sm {
			sm = arr[i]
			index = i
		}
	}
	return index
}
