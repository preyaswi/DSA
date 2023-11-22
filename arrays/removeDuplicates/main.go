package main

import "fmt"

func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    var new []int
    new = append(new, nums[0])
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[i-1] {
            new = append(new, nums[i])
        }
    }
    copy(nums, new)
    length := len(new)
    return length
}

func main() {
    input := []int{1, 1, 2, 2, 3, 4, 5, 5}
    
    fmt.Println("Original array:", input)
    
    newLength := removeDuplicates(input)
    
    fmt.Println("Array after removing duplicates:", input[:newLength])
    fmt.Println("New length:", newLength)
}
