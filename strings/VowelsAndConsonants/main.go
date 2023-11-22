package main


import (
    "fmt"
    "strings"
    "unicode"
)


func main() {
    input := "hhELLO  WORLdd"
    Vowels, Consonants := CountVowelsAndConsonants(input)
    fmt.Printf("the string %s\nthe vowels %d\nthe consonants %d\n", input, Vowels, Consonants)
}
func CountVowelsAndConsonants(input string) (int, int) {
    Vowels := "aeiouAEIOU"
    isvowelsCount, isConsonantsCount := 0, 0
    for _, char := range input {
        if char == ' ' {
            continue
        }
        if strings.ContainsRune(Vowels, char) {


            isvowelsCount++
        } else if unicode.IsLetter(char) {
            isConsonantsCount++
        }


    }
    return isvowelsCount, isConsonantsCount
}


