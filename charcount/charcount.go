package charcount

import "fmt"

type CharCount map[rune]int

func New() *CharCount {
	return &CharCount{}
}

func (c *CharCount) Add(char rune) {
	if count, ok := (*c)[char]; ok {
		count++
		(*c)[char] = count
	} else {
		(*c)[char] = 1
	}
}

func (c *CharCount) Print() {
	for v, count := range *c {
		fmt.Println("Char", string(v), "count", count)
	}
}
