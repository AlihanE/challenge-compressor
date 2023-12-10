package main

import (
	"bufio"
	"fmt"
	"io"
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
	s3 := bufio.NewReader(f3)
	s := huffman.NewSorter()
	line, _, err := s3.ReadLine()
	if err != nil {
		panic(err)
	}
	s.Parse(string(line))

	w := writer.New(filename, "txt", "-dec")
	defer w.Close()

	tree := huffman.NewTree(s)
	tree.BuildTree()
	code := []byte{}
	for {
		b, err := s3.ReadByte()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		code = append(code, b)
		node := huffman.GetLeafByCode(tree.Root, code, 0)
		if node.IsLeaf() {
			code = []byte{}
			w.Writer.WriteRune(node.Char())
		}
	}
	w.Writer.Flush()
	fmt.Println()
}

func encode(filename string) {
	f := openFile(filename)
	s := bufio.NewReader(f)

	counter := charcount.New()
	for {
		r, _, err := s.ReadRune()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		counter.Add(r)
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

	nn, err := w.Writer.Write([]byte(sorterString + "\r\n"))
	if err != nil {
		panic(err)
	}

	fmt.Println("nn", nn)

	f2 := openFile(filename)
	defer f2.Close()
	s2 := bufio.NewReader(f2)
	for {
		r, _, err := s2.ReadRune()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		val := table[r]
		for _, c := range val {
			err = w.Writer.WriteByte(c)
			if err != nil {
				panic(err)
			}
		}
	}
	w.Writer.Flush()
}

func openFile(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return f
}
