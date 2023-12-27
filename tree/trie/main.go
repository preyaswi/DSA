package main

import (
	"fmt"
)

const size = 26

type node struct {
	children [size]*node
	isEnd    bool
}
type trie struct {
	root *node
}

func main() {
	t := trie{&node{}}
	t.insert("preyaswi")
	t.insert("preya")
	b := t.search("preya")
	fmt.Println(b)
	c := t.search("s")
	fmt.Println(c)
	t.delete("preya")
	d := t.search("preya")
	fmt.Println(d)
}
func (t *trie) insert(s string) {
	currentIdx := t.root
	for _, v := range s {
		charIdx := v - 'a'
		if currentIdx.children[charIdx] == nil {
			currentIdx.children[charIdx] = &node{}
		}
		currentIdx = currentIdx.children[charIdx]
	}
	currentIdx.isEnd = true
}
func (t *trie) search(s string) bool {
	currentNode := t.root
	for i := 0; i < len(s); i++ {
		charIdx := s[i] - 'a'
		if currentNode.children[charIdx] == nil {
			return false
		}
		currentNode = currentNode.children[charIdx]
	}
	return currentNode.isEnd
}
func (t *trie) delete(s string) {
	t.root = t.deleteKey(t.root, s, 0)
}
func (t *trie) deleteKey(currentNode *node, s string, depth int) *node {
	if currentNode == nil {
		return nil
	}
	if depth == len(s) {
		currentNode.isEnd = false
		if isempty(currentNode) {
			return nil
		}
	}
	if depth<len(s){

		charIdx := s[depth] - 'a'
		currentNode.children[charIdx] = t.deleteKey(currentNode.children[charIdx], s, depth+1)
	}
	if isempty(currentNode) && !currentNode.isEnd {
		return nil
	}
	return currentNode
}
func isempty(currentNode *node) bool {
	for i := 0; i < size; i++ {
		if currentNode.children[i] != nil {
			return false
		}
	}
	return true
}
