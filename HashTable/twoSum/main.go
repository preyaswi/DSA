package main

import "fmt"

func main() {
	array := []int{2, 5, 9, 4, 7, 3}
	target := 8
	newmap := twoSum(array, target)
	fmt.Printf("the target %d is the sum of %d", target, newmap)
}
func twoSum(array []int, target int) []int {
	newtable := make(map[int]int)
	for i, num := range array {
		result := target - num
		if j, found := newtable[result]; found {
			return []int{j, i}
		}
		newtable[num] = i
	}
	return []int{}
}
