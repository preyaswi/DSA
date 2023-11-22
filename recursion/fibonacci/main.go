package main


import "fmt"


func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}


func main() {
    result := fibonacci(6)
    fmt.Println("Fibonacci of 6th position:", result)
}
