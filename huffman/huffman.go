package huffman

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

type BuildState string

const (
	First  BuildState = "first"
	Second BuildState = "second"
	New    BuildState = "new"
)

type Tree struct {
	state BuildState
	Root  *Node
	count int
}

func NewTree() *Tree {
	return &Tree{
		state: New,
	}
}

func (t *Tree) AddNode(node *Node) {
	if t.count == 0 {
		t.Root = node
		t.count++
		t.state = First
	} else {
		switch t.state {
		case First:

		}
	}
}

type Sorter []*Node

func NewSorter() Sorter {
	return Sorter{}
}

func (s Sorter) Len() int {
	return len(s)
}

func (s Sorter) Less(i, j int) bool {
	return s[i].Weight() < s[j].Weight()
}

func (s Sorter) Swap(i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}
