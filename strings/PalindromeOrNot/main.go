package main


import (
    "fmt"
    "strings"
    "unicode"
)


func main() {
    input := "A man, a plan, a canal, Panama"
    palindrom := IsPalindrom(input)
    fmt.Printf("the given string %s is %t\n", input, palindrom)
}
func IsPalindrom(input string) bool {
    input = strings.ToLower(input)
    input = strings.ReplaceAll(input, " ", "")
    var cleanedInput strings.Builder
    for _, char := range input {
        if unicode.IsLetter(char) || unicode.IsDigit(char) {
            cleanedInput.WriteRune(char)
        }


    }
    cleanedString := cleanedInput.String()
    length := len(cleanedString)
    for i := 0; i < length/2; i++ {
        if cleanedString[i] != cleanedString[length-i-1] {


            return false
        }


    }
    return true
}
