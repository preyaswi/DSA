package main


import "fmt"


func linearSearch(arr []int, target int) []int {
    var ar []int


    for i, element := range arr {
        if element == target {
            ar = append(ar, i)


        }
    }
    return ar
}


func main() {
    array := []int{16, 5, 8, 12, 16, 23, 16}
    targetElement := 16
    resultArray := linearSearch(array, targetElement)
    if len(resultArray) > 0 {
        fmt.Println("Element found in these postions ", resultArray)
    } else {
        fmt.Println("Elements not found")
    }


}
