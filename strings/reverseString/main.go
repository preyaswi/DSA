package main


import "fmt"


func main() {
    input := "hello"


    reverstring := Reverse(input)
    fmt.Printf("original %s\nRevesed %s\n", input, reverstring)


}
func Reverse(input string) string {
    inputinrune := []rune(input)
    length := len(inputinrune)
    for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
        inputinrune[i], inputinrune[j] = inputinrune[j], inputinrune[i]


    }
    inputstring := string(inputinrune)
    return inputstring


}
