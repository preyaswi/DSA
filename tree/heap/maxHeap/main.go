package main

import "fmt"

type maxheap struct {
	arr []int
}

func main() {
	arr := []int{11, 2, 7, 4,12, 5, 9,10}
	h := &maxheap{}
	h.buildHeap(arr)
	fmt.Println("original array:")
	display(h.arr)
h.heapSort()
fmt.Println("sorted array:")
display(h.arr)

}
func(h *maxheap) buildHeap(arr []int)  {
	h.arr=arr
	for i:=parent(len(h.arr)-1);i>=0;i--{
h.shiftdown(i)
	}
}
func parent(i int)int  {
	return (i-1)/2
}
func (h *maxheap)shiftdown(currentIdx int)  {
	endIdx:=len(h.arr)-1
	leftIdx:=leftchild(currentIdx)
	for leftIdx<=endIdx{
		var IdxToSwap int
		rightIdx:=rightchild(currentIdx)
		if rightIdx<=endIdx&&h.arr[rightIdx]>h.arr[leftIdx]{
			IdxToSwap=rightIdx
		}else {
			IdxToSwap=leftIdx
		}
		if h.arr[IdxToSwap]>h.arr[currentIdx]{
			swap(h.arr,IdxToSwap,currentIdx)
			currentIdx=IdxToSwap
			leftIdx=leftchild(currentIdx)
		}else {
			return
		}
	}
}
func leftchild(i int)int  {
	return 2*i+1
}
func rightchild(i int) int {
	return 2*i+2
}
func swap(arr []int,a int,b int)  {
	arr[a],arr[b]=arr[b],arr[a]
}
func display(arr []int)  {
	for _,v:=range arr{
		fmt.Print(v," ")
	}
	fmt.Println()
}
func (h *maxheap)heapSort()[]int  {
	for i:=len(h.arr)-1;i>=0;i--{
		h.arr[0],h.arr[i]=h.arr[i],h.arr[0]
		h.shiftdown(0)
	}
	return h.arr
}