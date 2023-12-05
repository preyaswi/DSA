package main

import "fmt"

func main() {
	arr := []int{9, 4, 2, 5, 2, 1, 9, 8, 1}
	fmt.Println("before sorting", arr)
	AsecndingOrder(arr)
	fmt.Println("after sorting", arr)

}
func AsecndingOrder(arr []int) {
	n := len(arr)
	var swap bool
	for i := 0; i < n-1; i++ {
		swap = false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swap = true
			}

		}
		if !swap {
			break
		}
	}
}
