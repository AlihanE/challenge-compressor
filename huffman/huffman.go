package huffman

import (
	"fmt"
)

type Node struct {
	isLeaf    bool
	weight    int
	char      rune
	LeftLeaf  *Node
	RightLeaf *Node
}

func NewNode(isLeaf bool, weight int, char rune) *Node {
	return &Node{
		isLeaf: isLeaf,
		weight: weight,
		char:   char,
	}
}

func (n *Node) IsLeaf() bool {
	return n.isLeaf
}

func (n *Node) Weight() int {
	return n.weight
}

func (n *Node) SetWeight(w int) {
	n.weight = w
}

func (n *Node) Char() rune {
	return n.char
}

type Sorter struct {
	arr []*Node
}

func NewSorter() Sorter {
	return Sorter{}
}

func (s *Sorter) Add(n *Node) {
	length := len(s.arr)
	if length == 0 {
		s.arr = append(s.arr, n)
	} else {
		if s.arr[length-1].Weight() > n.Weight() {
			s.arr = append([]*Node{n}, s.arr...)
			fmt.Println(s.arr)
		} else {
			s.arr = append(s.arr, n)
		}
	}
}

func (s *Sorter) SortedArray() []*Node {
	return s.arr
}
