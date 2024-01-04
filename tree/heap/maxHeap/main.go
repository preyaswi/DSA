package main

import "fmt"

type maxHeap struct {
	arr []int
}

func (h *maxHeap) insert(data int) {
	h.arr = append(h.arr, data)
	h.buildHeap()
}

func (h *maxHeap) buildHeap() {
	n := len(h.arr)
	for i := n/2 - 1; i >= 0; i-- {
		h.shiftDown(i)
	}
}

func (h *maxHeap) shiftDown(currentIdx int) {
	endIdx := len(h.arr) - 1

	for {
		leftIdx := leftChild(currentIdx)
		if leftIdx > endIdx {
			break
		}

		var idxToSwap int
		
		rightIdx := rightChild(currentIdx)
		if rightIdx <= endIdx && h.arr[rightIdx] > h.arr[leftIdx] {
			idxToSwap = rightIdx
		} else {
			idxToSwap = leftIdx
		}

		if h.arr[idxToSwap] > h.arr[currentIdx] {
			swap(h.arr, idxToSwap, currentIdx)
			currentIdx = idxToSwap
		} else {
			break
		}
	}
}

func (h *maxHeap) delete() int {
	if len(h.arr) == 0 {
		return -1
	}
	maxValue := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.shiftDown(0)

	return maxValue
}

func (h *maxHeap) heapSort() []int {
	n := len(h.arr)

	for i := n - 1; i > 0; i-- {
		h.arr[0], h.arr[i] = h.arr[i], h.arr[0]
		h.shiftDown(0)
	}

	return h.arr
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

func parent(i int) int {
	return (i - 1) / 2
}

func swap(arr []int, a int, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func display(arr []int) {
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println(" ")
}

func main() {
	var arr = []int{11, 2, 7, 4, 12, 5, 9, 10}
	h := &maxHeap{}

	for _, val := range arr {
		h.insert(val)
	}

	fmt.Println("Original array:")
	display(h.arr)

	h.heapSort()

	fmt.Println("Sorted array:")
	display(h.arr)
}