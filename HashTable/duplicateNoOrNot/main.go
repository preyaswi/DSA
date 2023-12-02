package main

import "fmt"

func main() {
	array := []int{2, 3, 3, 5, 3, 6, 5, 8}
	fmt.Println("array before deleting the duplicate", array)
	newArray := DuplicateDelete(array)
	fmt.Println("array after deleting the duplicate", newArray)

}
func DuplicateDelete(array []int) []int {
	seen := make(map[int]bool)
	var newArray []int
	for _, v := range array {
		if !seen[v] {
			newArray = append(newArray, v)
			seen[v] = true
		}
	}
	return newArray
}
