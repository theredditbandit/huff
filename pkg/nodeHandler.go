package pkg

import (
	"sort"
)

type Node struct {
	Freq  int
	Val   string
	Left  *Node
	Right *Node
	Up    *Node
}

type SortedNodes []Node

func (s SortedNodes) Len() int {
	return len(s)
}

func (s *SortedNodes) Add(n Node) {
	*s = append(*s, n)
	sort.Slice(*s, func(i, j int) bool {
		return (*s)[i].Freq < (*s)[j].Freq
	})
}

func (s *SortedNodes) Pop() Node {
	n := (*s)[0]
	*s = (*s)[1:]
	return n
}
