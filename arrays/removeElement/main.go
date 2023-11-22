package main

import "fmt"

func removeElement(nums []int, val int) int {
    var newindex int = 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[newindex] = nums[i]
            newindex++
        }
    }
    return newindex
}

func main() {
    input := []int{3, 2, 2, 3, 4, 5, 6}
    valToRemove := 3
    
    fmt.Println("Original array:", input)
    
    newLength := removeElement(input, valToRemove)
    
    fmt.Println("Array after removing element:", input[:newLength])
    fmt.Println("New length:", newLength)
}
