package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/AlihanE/challenge-compressor/bit"
	"github.com/AlihanE/challenge-compressor/charcount"
	"github.com/AlihanE/challenge-compressor/huffman"
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

	line, _, err = s3.ReadLine()
	if err != nil {
		panic(err)
	}

	charCount, err := strconv.Atoi(string(line))
	if err != nil {
		panic(err)
	}

	nameWithoutExt := GetFileNameWithoutExt(filename)
	f, err := os.Create("./" + nameWithoutExt + "-dec" + "." + "pee")
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)

	tree := huffman.NewTree(s)
	tree.BuildTree()
	code := []byte{}
	cont := true
	for cont {
		b, err := s3.ReadByte()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		for _, bit := range bit.GetBitArray(b) {
			code = append(code, bit)
			node := huffman.GetLeafByCode(tree.Root, code, 0)
			if node.IsLeaf() {
				charCount--
				code = []byte{}
				w.WriteRune(node.Char())
				if charCount == 0 {
					cont = false
					break
				}
			}
		}

	}
	w.Flush()
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

	nameWithoutExt := GetFileNameWithoutExt(filename)
	f, err := os.Create("./" + nameWithoutExt + "." + "pee")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	nn, err := w.Write([]byte(sorterString + "\r\n"))
	if err != nil {
		panic(err)
	}

	fmt.Println("nn", nn)

	nn, err = w.Write([]byte(strconv.Itoa(counter.TotalChars()) + "\r\n"))
	if err != nil {
		panic(err)
	}

	w.Flush()

	fmt.Println("nn", nn)

	f2 := openFile(filename)
	defer f2.Close()
	s2 := bufio.NewReader(f2)

	b := bit.NewBitWriter(f)

	for {
		r, _, err := s2.ReadRune()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		val := table[rune(r)]

		b.AddBits(val)
	}
	b.Close()
}

func openFile(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return f
}

func GetFileNameWithoutExt(name string) string {
	return name[:strings.Index(name, ".")]
}
