package main


import "fmt"


func linearSearch(arr []int, target int) int {
    for i, element := range arr {
        if element == target {
            return i
        }
    }
    return -1
}


func main() {
    array := []int{2, 5, 8, 12, 16, 23}
    targetElement := 16


    result := linearSearch(array, targetElement)


    if result != -1 {
        fmt.Printf("Element %d found at index %d\n", targetElement, result)
    } else {
        fmt.Printf("Element %d not found in the array\n", targetElement)
    }
}
