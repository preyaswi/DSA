package main

import "fmt"

type vertex struct {
	val      int
	adjacent []*vertex
}
type graph struct {
	vetices []*vertex
}

func (g *graph) addvertex(val int) {
	if !contains(g.vetices, val) {
		g.vetices = append(g.vetices, &vertex{val: val})
	}
}
func contains(vertex []*vertex, val int) bool {
	for _, v := range vertex {
		if v.val == val {

			return true
		}
	}
	return false
}
func (g *graph) addedges(from, to int) {
	fromvertex := getvetex(g.vetices, from)
	tovertex := getvetex(g.vetices, to)
	if fromvertex == nil || tovertex == nil {
		fmt.Println("vertex not found")
		return
	}
	fromvertex.adjacent = append(fromvertex.adjacent, tovertex)
	tovertex.adjacent = append(tovertex.adjacent, fromvertex)
}
func getvetex(vertex []*vertex, val int) *vertex {
	for _, v := range vertex {
		if v.val == val {

			return v
		}
	}
	return nil
}
func (g *graph) display() {
	for _, v := range g.vetices {
		fmt.Print("the vertex ", v.val, ":")
		for _, neighbor := range v.adjacent {
			fmt.Print(" ", neighbor.val, " ")
		}
		fmt.Println()
	}
}

// bfs traversal
type queue struct {
	arr []*vertex
}

func (g *graph) bfs(key int) {
	q := &queue{}
	ischecked := make(map[int]bool)
	stavertex := getvetex(g.vetices, key)
	if stavertex == nil {
		fmt.Println("vertex not found")
		return
	}
	ischecked[stavertex.val] = true
	q.arr = append(q.arr, stavertex)
	for len(q.arr) != 0 {
		vertex := q.arr[0]
		q.arr=q.arr[1:]
		fmt.Print(vertex.val, " ")
		for _, neighbors := range vertex.adjacent {
			if !ischecked[neighbors.val] {
				ischecked[neighbors.val] = true
				q.arr = append(q.arr, neighbors)
			}
		}
	}
	fmt.Println()

}
//dfs traversal
type stack struct{
	arr []*vertex
}
func (g *graph)dfs(key int)  {
	s:=stack{}
	ischecked:=make(map[int]bool)
	startvertex:=getvetex(g.vetices,key)
	if startvertex==nil{
		fmt.Println("vertex not found")
	}
	ischecked[key]=true
	s.arr=append(s.arr, startvertex)
	for len(s.arr)>0{
		vertex:=s.arr[len(s.arr)-1]
		s.arr=s.arr[:len(s.arr)-1]
		fmt.Print(" ",vertex.val," ")
		for _,neighbor:=range vertex.adjacent{
			if !ischecked[neighbor.val]{
				ischecked[neighbor.val]=true
				s.arr=append(s.arr, neighbor)
			}
		}
	}
	fmt.Println()
}

func main() {
	g := &graph{}
	g.addvertex(1)
	g.addvertex(2)
	g.addvertex(3)
	g.addvertex(4)
	g.addvertex(5)
	g.addvertex(6)

	g.addedges(1, 3)
	g.addedges(1, 2)
	g.addedges(1, 5)
	g.addedges(2, 3)
	g.addedges(5, 3)
	g.addedges(6, 2)
	g.addedges(4, 6)
	g.display()
	fmt.Println("bfs")
	g.bfs(1)
	fmt.Println("dfs")
	g.dfs(1)
}
