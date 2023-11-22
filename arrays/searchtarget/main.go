package main

import "fmt"

func searchInsert(nums []int, target int) int {
    var mid int
    low, high := 0, len(nums)-1
    for low <= high {
        mid = (low + high) / 2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            low = mid + 1
        } else if nums[mid] > target {
            high = mid - 1
        }
    }
    return low
}

func main() {
    sortedArray := []int{1, 3, 5, 6}
    target1 := 5
    target2 := 2
    target3 := 7
    target4 := 0
    
    result1 := searchInsert(sortedArray, target1)
    result2 := searchInsert(sortedArray, target2)
    result3 := searchInsert(sortedArray, target3)
    result4 := searchInsert(sortedArray, target4)
    
    fmt.Printf("Target %d should be inserted at index %d\n", target1, result1)
    fmt.Printf("Target %d should be inserted at index %d\n", target2, result2)
    fmt.Printf("Target %d should be inserted at index %d\n", target3, result3)
    fmt.Printf("Target %d should be inserted at index %d\n", target4, result4)
}
