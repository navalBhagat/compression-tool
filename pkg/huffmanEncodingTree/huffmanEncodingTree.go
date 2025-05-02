package huffmanEncodingTree

import (
	"fmt"
	"sort"
	"strings"
)

type HuffBaseNode interface {
	IsLeaf() bool
	Weight() int
}

type HuffLeafNode struct {
	Element    string
	NodeWeight int
}

func (h HuffLeafNode) IsLeaf() bool {
	return true
}

func (h HuffLeafNode) Weight() int {
	return h.NodeWeight
}

func (h HuffLeafNode) Value() string {
	return h.Element
}

type HuffInternalNode struct {
	NodeWeight int
	Left       HuffBaseNode
	Right      HuffBaseNode
}

func (h HuffInternalNode) IsLeaf() bool {
	return false
}

func (h HuffInternalNode) Weight() int {
	return h.NodeWeight
}

func (h HuffInternalNode) LeftChild() HuffBaseNode {
	return h.Left
}

func (h HuffInternalNode) RightChild() HuffBaseNode {
	return h.Right
}

type HuffTree struct {
	Root HuffBaseNode
}

func NewHuffTree(el string, wt int) HuffTree {
	return HuffTree{Root: HuffLeafNode{Element: el, NodeWeight: wt}}
}

func NewHuffTreeWithChildren(l HuffBaseNode, r HuffBaseNode, wt int) HuffTree {
	return HuffTree{Root: HuffInternalNode{NodeWeight: wt, Left: l, Right: r}}
}

func (h HuffTree) RootNode() HuffBaseNode {
	return h.Root
}

func (h HuffTree) Weight() int {
	return h.Root.Weight()
}

func BuildTree(freqMap map[string]int) HuffTree {
	var trees []HuffTree
	for el, wt := range freqMap {
		trees = append(trees, NewHuffTree(el, wt))
	}
	for len(trees) > 1 {
		sort.Slice(trees, func(i, j int) bool {
			return trees[i].Weight() < trees[j].Weight()
		})
		trees = append(trees, CombineTrees(trees[0], trees[1]))
		trees = trees[2:]
	}
	return trees[0]
}

func CombineTrees(t1 HuffTree, t2 HuffTree) HuffTree {
	wt := t1.Weight() + t2.Weight()
	return NewHuffTreeWithChildren(t1.RootNode(), t2.RootNode(), wt)
}

func PrettyPrint(node HuffBaseNode, prefix string) {
	if node.IsLeaf() {
		leaf := node.(HuffLeafNode)
		fmt.Printf("%sLeaf: '%s' (Weight: %d)\n", prefix, leaf.Element, leaf.NodeWeight)
	} else {
		internal := node.(HuffInternalNode)
		fmt.Printf("%sInternal (Weight: %d)\n", prefix, internal.NodeWeight)
		PrettyPrint(internal.Left, prefix+strings.Repeat(" ", 4))
		PrettyPrint(internal.Right, prefix+strings.Repeat(" ", 4))
	}
}

func PrefixTable(tree HuffTree) map[string]string {
	prefixTable := make(map[string]string)
	var buildTable func(node HuffBaseNode, prefix string)
	buildTable = func(node HuffBaseNode, prefix string) {
		if node.IsLeaf() {
			leaf := node.(HuffLeafNode)
			prefixTable[leaf.Element] = prefix
		} else {
			internal := node.(HuffInternalNode)
			buildTable(internal.LeftChild(), prefix+"0")
			buildTable(internal.RightChild(), prefix+"1")
		}
	}
	buildTable(tree.RootNode(), "")
	return prefixTable
}
