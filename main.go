package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/AlihanE/challenge-compressor/charcount"
	"github.com/AlihanE/challenge-compressor/huffman"
)

func main() {
	if len(os.Args) < 2 {
		panic("File name needed")
	}

	f := openFile(os.Args[1])
	s := bufio.NewScanner(f)

	counter := charcount.New()
	for s.Scan() {
		for _, b := range s.Bytes() {
			counter.Add(rune(b))
		}
	}

	sorter := huffman.NewSorter()
	fmt.Println("Counter length", len(*counter))
	for v, c := range *counter {
		n := huffman.NewNode(true, c, v)
		sorter.AddNode(n)
	}

	sort.Sort(sorter)

	sorter.Print()

	tree := huffman.NewTree(sorter)

	table := tree.GetTable()
	for k, v := range table {
		fmt.Println("key", string(k), "val", v)
	}

	fmt.Println(tree.Root.String())
}

func openFile(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return f
}
