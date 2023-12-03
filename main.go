package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/AlihanE/challenge-compressor/charcount"
	"github.com/AlihanE/challenge-compressor/huffman"
	"github.com/AlihanE/challenge-compressor/writer"
)

func main() {
	if len(os.Args) < 3 {
		panic("File name needed")
	}

	if os.Args[1] == "-e" {
		encode(os.Args[2])
	} else if os.Args[1] == "-d" {
		decode(os.Args[2])
	}
}

func decode(filename string) {
	f3 := openFile(filename)
	s3 := bufio.NewScanner(f3)
	s := huffman.NewSorter()
	if s3.Scan() {
		tt := s3.Text()
		s.Parse(string(tt))
	}

	w := writer.New(filename, "txt", "-dec")
	defer w.Close()

	tree := huffman.NewTree(s)
	tree.BuildTree()
	for s3.Scan() {
		code := []byte{}
		for _, b := range s3.Bytes() {
			code = append(code, b)
			node := huffman.GetLeafByCode(tree.Root, code, 0)
			if node.IsLeaf() {
				code = []byte{}
				w.Write([]byte{byte(node.Char())})
			}
		}
	}
	fmt.Println()
}

func encode(filename string) {
	f := openFile(filename)
	s := bufio.NewScanner(f)

	counter := charcount.New()
	for s.Scan() {
		for _, b := range s.Bytes() {
			counter.Add(rune(b))
		}
	}
	f.Close()
	sorter := huffman.NewSorter()
	fmt.Println("Counter length", len(*counter))
	for v, c := range *counter {
		n := huffman.NewNode(true, c, v)
		sorter.AddNode(n)
	}

	sort.Sort(sorter)

	sorter.Print()
	sorterString := sorter.String()
	tree := huffman.NewTree(sorter)

	table := tree.GetTable()
	for k, v := range table {
		fmt.Println("key", string(k), "val", v)
	}

	w := writer.New(filename, "pee", "")
	defer w.Close()

	_, err := w.Write([]byte(sorterString + "\r\n"))
	if err != nil {
		panic(err)
	}

	f2 := openFile(filename)
	defer f2.Close()
	s2 := bufio.NewScanner(f2)
	for s2.Scan() {
		for _, b := range s2.Bytes() {
			val := table[rune(b)]
			_, err := w.Write(val)
			if err != nil {
				panic(err)
			}
		}
	}
}

func openFile(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return f
}
