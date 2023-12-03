package huffman

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
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

func (n *Node) GetNodeByCode(code []byte) *Node {
	node := n
	for _, v := range code {
		if v > 0 {
			node = node.RightLeaf
		} else {
			node = node.LeftLeaf
		}
	}

	return node
}

func (n *Node) String() string {
	left := ""
	if n.LeftLeaf != nil {
		left = n.LeftLeaf.String()
	}

	right := ""
	if n.RightLeaf != nil {
		right = n.RightLeaf.String()
	}
	return fmt.Sprintf("(%t, %d, %d, %s, %s)", n.isLeaf, n.weight, n.char, left, right)
}

type Tree struct {
	Root   *Node
	count  int
	sorter *Sorter
}

func NewTree(nodes *Sorter) *Tree {
	return &Tree{
		sorter: nodes,
	}
}

func (t *Tree) BuildTree() {
	has := len(t.sorter.nodes) > 1
	for has {
		first := t.sorter.GetFirst()
		second := t.sorter.GetFirst()
		if second == nil {
			t.Root = first
			has = false
			return
		}

		newNode := NewNode(false, first.Weight()+second.Weight(), -1)
		if first.Weight() > second.Weight() {
			newNode.LeftLeaf = second
			newNode.RightLeaf = first
		} else {
			newNode.LeftLeaf = first
			newNode.RightLeaf = second
		}
		t.sorter.AddNode(newNode)
		sort.Sort(t.sorter)
	}
}

// func (t *Tree) printNodes() {
// 	for _, v := range t.sorter.nodes {
// 		fmt.Printf("%d - %s", v.Weight(), string(v.Char()))
// 	}
// 	fmt.Println()
// }

func (t *Tree) Count() int {
	return t.count
}

func (t *Tree) GetTable() map[rune][]byte {
	table := make(map[rune][]byte)
	code := []byte{}

	t.BuildTree()

	buildPrefixTable(t.Root, code, table)

	return table
}

func buildPrefixTable(node *Node, code []byte, table map[rune][]byte) {
	if node == nil {
		return
	}

	if node.isLeaf {
		dst := make([]byte, len(code))
		copy(dst, code)
		table[node.char] = dst
	} else {
		left := append(code, 0)
		buildPrefixTable(node.LeftLeaf, left, table)
		right := append(code, 1)
		buildPrefixTable(node.RightLeaf, right, table)
	}
}

func GetLeafByCode(node *Node, code []byte, bit int) *Node {
	if node == nil {
		return nil
	}

	var res *Node
	if code[bit] == 0 {
		res = node.LeftLeaf
	} else if code[bit] == 1 {
		res = node.RightLeaf
	}

	if len(code) == bit+1 {
		return res
	}

	return GetLeafByCode(res, code, bit+1)
}

type Sorter struct {
	nodes []*Node
}

func NewSorter() *Sorter {
	return &Sorter{}
}

func (s *Sorter) Len() int {
	return len(s.nodes)
}

func (s *Sorter) Less(i, j int) bool {
	return s.nodes[i].Weight() < s.nodes[j].Weight()
}

func (s *Sorter) Swap(i, j int) {
	tmp := s.nodes[i]
	s.nodes[i] = s.nodes[j]
	s.nodes[j] = tmp
}

func (s *Sorter) GetFirst() *Node {
	var first *Node
	if len(s.nodes) > 0 {
		first = s.nodes[0]
		s.nodes = s.nodes[1:]
	}
	return first
}

func (s *Sorter) AddNode(node *Node) {
	s.nodes = append(s.nodes, node)
}

func (s *Sorter) Print() {
	for _, v := range s.nodes {
		fmt.Println("literal", string(v.Char()), "weight", v.Weight())
	}
}

func (s *Sorter) String() string {
	sl := []string{}
	for _, v := range s.nodes {
		sl = append(sl, fmt.Sprintf("%d,%d", v.char, v.weight))
	}
	return strings.Join(sl, ";")
}

func (s *Sorter) Parse(input string) {
	arr := strings.Split(input, ";")
	for _, v := range arr {
		subArr := strings.Split(v, ",")
		weight, err := strconv.Atoi(subArr[1])
		if err != nil {
			panic(err)
		}
		char, err := strconv.Atoi(subArr[0])
		if err != nil {
			panic(err)
		}
		s.AddNode(NewNode(true, weight, rune(char)))
	}
}

func PrintTree(node *Node, side string) {
	if node == nil {
		return
	}

	fmt.Println("leaf", node.IsLeaf(), "weight", node.Weight(), "side", side, "literal", string(node.char))
	PrintTree(node.LeftLeaf, "left")
	PrintTree(node.RightLeaf, "right")
}

func (t *Tree) PrintTreeByLevel() {
	levels := make(map[int][]*Node)

	getLevelsAndNodes(t.Root, 0, levels)
	for i := 0; i < len(levels); i++ {
		for _, v := range levels[i] {

			fmt.Printf("(%d %s)      ", v.Weight(), string(v.Char()))
		}
		fmt.Println("")
		fmt.Println("-------------------------------")
	}
}

func getLevelsAndNodes(node *Node, level int, levels map[int][]*Node) {
	if node == nil {
		return
	}

	levels[level] = append(levels[level], node)

	getLevelsAndNodes(node.LeftLeaf, level+1, levels)
	getLevelsAndNodes(node.RightLeaf, level+1, levels)
}
