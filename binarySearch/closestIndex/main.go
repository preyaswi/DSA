package main
import (
    "fmt"
    "math"
)


func main() {
    array := []int{2, 5, 8, 12, 16, 23, 28, 32, 40, 50}
    targetValue := 90
    closestIndex := Closest(array, targetValue)
    if closestIndex != -1 {
        fmt.Printf("Closest element to %d is %d at index %d\n", targetValue, array[closestIndex], closestIndex)
    } else {
        fmt.Println("Array is empty")
    }
}
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
func Closest(arr []int, target int) int {
    low, high := 0, len(arr)-1
    clos := math.MaxInt64
    closestIndex := -1


    for low <= high {
        mid := (low + high) / 2


        if abs(arr[mid]-target) < abs(clos-target) {
            closestIndex = mid
            clos = arr[mid]


        }


        if arr[mid] == target {
            return mid
        }
        if arr[mid] < target {
            low = mid + 1
        }
        if arr[mid] > target {
            high = mid - 1
        }
    }
    return closestIndex


}
