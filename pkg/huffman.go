// Description: Huffman encoder implementation
package pkg

import (
	"sort"
)

type Node struct {
	Freq  int
	Val   string
	Left  *Node
	Right *Node
}

func BuildHuffmanTree(stats map[string]int) {
	nodes := GenNodes(stats)
	_ = nodes
}

func GenNodes(stats map[string]int) []Node {
	var treeNodes []Node
	for k, v := range stats {
		n := Node{
			Freq: v,
			Val:  k,
		}
		treeNodes = append(treeNodes, n)
	}
	sort.Slice(treeNodes, func(i, j int) bool {
		return treeNodes[i].Freq < treeNodes[j].Freq
	})
	return treeNodes
}
