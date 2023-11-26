// sorted by the score of students in ascending order
package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Score float64
}

func main() {

	Students := []Person{
		{"preya", 20, 89.8},
		{"sravan", 23, 95},
		{"pranal", 26, 98},
		{"sreya", 21, 90},
	}
	fmt.Println("the unsorted struct", Students)

	SortByScore(Students)
	fmt.Println("sorted by the score of students", Students)
}
func SortByScore(students []Person) {
	n := len(students)
	for i := 0; i < n-1; i++ {
		midInd := i
		for j := i + 1; j < n; j++ {
			if students[j].Score < students[midInd].Score {
				midInd = j
			}
		}
		students[i], students[midInd] = students[midInd], students[i]
	}
}
