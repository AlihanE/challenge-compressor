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
		sorter = append(sorter, n)
	}

	sort.Sort(sorter)

	for _, v := range sorter {
		fmt.Print((*v).Char(), " ", (*v).Weight(), " | ")
	}
}

func openFile(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return f
}
