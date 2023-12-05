// sorting a struct of slice by name in ascending order
package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Score int
}

func main() {
	students := []Person{
		{"sravan", 34, 89},
		{"preya", 23, 133},
		{"sreyaswi", 23, 483},
	}
	fmt.Println("before sorting by name", students)
	SortByName(students)
	fmt.Println("after sorting by name", students)

}
func SortByName(students []Person) {
	n := len(students)
	for i := 0; i < n-1; i++ {
var swap bool
		for j := 0; j < n-i-1; j++ {
			swap=false
			if students[j].Name > students[j+1].Name {
				students[j], students[j+1] = students[j+1], students[j]
				swap=true
			}
		}
		if !swap{
			break
		}
	}
}
