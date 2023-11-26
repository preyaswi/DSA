// sort the array in ascending order by length of the string
package main

import "fmt"

func main() {
	arr := []string{"preya", "sreyaswi", "pranal", "prakashan", "sravan", "latha"}
	fmt.Println("unsorted array", arr)
	SortByLength(arr)
	fmt.Println("array sorted by length", arr)
}
func SortByLength(arr []string) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minLen := i
		for j := i + 1; j < n; j++ {
			if len(arr[j]) < len(arr[minLen]) {
				minLen = j
			}
		}
		arr[i], arr[minLen] = arr[minLen], arr[i]
	}
}
