package main

import "fmt"

type minheap struct {
	array []int
}

func main() {
	arr := []int{8, 4, 2, 5, 22, 7, 4, 2, 44}
	h := &minheap{}
	for _, v := range arr {
		h.insert(v)
	}
	h.build() 
	h.display()
}

func (h *minheap) insert(data int) {
	h.array = append(h.array, data)
}

func (h *minheap) build() {
	for i := parent(len(h.array) - 1); i >= 0; i-- {
		h.shiftDown(i)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func (h *minheap) shiftDown(currentIdx int) {
	endIdx := len(h.array) - 1
	var idxToSwap int

	for {
		leftIdx := leftchild(currentIdx)
		rightIdx := rightChild(currentIdx)

		if leftIdx <= endIdx {
			if rightIdx <= endIdx && h.array[leftIdx] < h.array[rightIdx] {
				idxToSwap = leftIdx
			} else {
				idxToSwap = rightIdx
			}

			if h.array[idxToSwap] < h.array[currentIdx] {
				swap(h.array, idxToSwap, currentIdx)
				currentIdx = idxToSwap
			} else {
				break
			}
		} else {
			break
		}
	}
}

func leftchild(i int) int {
	return (2 * i) + 1
}

func rightChild(i int) int {
	return (2 * i) + 2
}

func swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func (h *minheap) display() {
	for _, v := range h.array {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
