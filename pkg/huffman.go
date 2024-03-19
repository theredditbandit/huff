// Description: Huffman encoder implementation
package pkg

type HuffmanTree struct {
	Root *Node
}

func BuildHuffmanTree(stats map[string]int) {
	nodes := GenNodes(stats)
	for nodes.Len() > 1 {
		left := nodes.Pop()
		right := nodes.Pop()
		parent := Node{
			Freq:  left.Freq + right.Freq,
			Left:  &left,
			Right: &right,
		}
		nodes.Add(parent)
	}
}

// GenNodes returns a slice of sorted nodes from the input map
func GenNodes(stats map[string]int) SortedNodes {
	var treeNodes SortedNodes
	for k, v := range stats {
		n := Node{
			Freq: v,
			Val:  k,
		}
		treeNodes.Add(n)
	}
	return treeNodes
}
