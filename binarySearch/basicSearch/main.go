package main

import "fmt"

func binarySearch(arr []int, target int) int {
    left, right := 0, len(arr)-1

    for left <= right {
        mid := left + (right-left)/2

        if arr[mid] == target {
            return mid 
        } else if arr[mid] < target {
            left = mid + 1 
        } else {
            right = mid - 1
        }
    }

    return -1 
}

func main() {
 
    array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    target1 := 5
    target2 := 11

    result1 := binarySearch(array, target1)
    result2 := binarySearch(array, target2)

    if result1 != -1 {
        fmt.Printf("Target %d found at index %d\n", target1, result1)
    } else {
        fmt.Printf("Target %d not found\n", target1)
    }

    if result2 != -1 {
        fmt.Printf("Target %d found at index %d\n", target2, result2)
    } else {
        fmt.Printf("Target %d not found\n", target2)
    }
}
