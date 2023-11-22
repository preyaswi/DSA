package main

import "fmt"

func mySqrt(x int) int {
    if x == 0 || x == 1 {
        return x
    }

    left, right := 1, x
    result := 0

    for left <= right {
        mid := left + (right-left)/2

        if mid*mid == x {
            return mid
        } else if mid*mid < x {
            result = mid
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return result
}

func main() {
    fmt.Println("Square root of 4:", mySqrt(4))
    fmt.Println("Square root of 8:", mySqrt(8))
    fmt.Println("Square root of 16:", mySqrt(16))
    fmt.Println("Square root of 25:", mySqrt(25))
}
