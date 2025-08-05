package huffmanEncodingTree

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
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
		sortHuffmanTreesByWeightAndTieBreak(trees)
		trees = append(trees, CombineTrees(trees[0], trees[1]))
		trees = trees[2:]
	}
	return trees[0]
}

func sortHuffmanTreesByWeightAndTieBreak(huffmanTrees []HuffTree) {
	sort.Slice(huffmanTrees, func(i, j int) bool {
		leftTree := huffmanTrees[i]
		rightTree := huffmanTrees[j]

		leftWeight := leftTree.Weight()
		rightWeight := rightTree.Weight()

		if leftWeight != rightWeight {
			return leftWeight < rightWeight
		}

		leftLeafNode, leftIsLeaf := leftTree.RootNode().(*HuffLeafNode)
		rightLeafNode, rightIsLeaf := rightTree.RootNode().(*HuffLeafNode)

		if leftIsLeaf && rightIsLeaf {
			return leftLeafNode.Element < rightLeafNode.Element
		}

		leftPointerString := fmt.Sprintf("%p", leftTree.RootNode())
		rightPointerString := fmt.Sprintf("%p", rightTree.RootNode())
		return leftPointerString < rightPointerString
	})
}

func CombineTrees(t1 HuffTree, t2 HuffTree) HuffTree {
	wt := t1.Weight() + t2.Weight()
	return NewHuffTreeWithChildren(t1.RootNode(), t2.RootNode(), wt)
}

func PrettyPrintTree(node HuffBaseNode, prefix string) {
	if node.IsLeaf() {
		leaf := node.(HuffLeafNode)
		fmt.Printf("%sLeaf: '%s' (Weight: %d)\n", prefix, leaf.Element, leaf.NodeWeight)
	} else {
		internal := node.(HuffInternalNode)
		fmt.Printf("%sInternal (Weight: %d)\n", prefix, internal.NodeWeight)
		PrettyPrintTree(internal.Left, prefix+strings.Repeat(" ", 4))
		PrettyPrintTree(internal.Right, prefix+strings.Repeat(" ", 4))
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

func SerializeTree(tree HuffTree) []byte {
	var buffer bytes.Buffer
	writeNode(tree.RootNode(), &buffer)
	return buffer.Bytes()
}

func writeNode(node HuffBaseNode, writer *bytes.Buffer) {
	if node.IsLeaf() {
		writer.WriteByte(1)
		leaf := node.(HuffLeafNode)
		r, _ := utf8.DecodeRuneInString(leaf.Element)
		writer.WriteRune(r)
	} else {
		writer.WriteByte(0)
		internal := node.(HuffInternalNode)
		writeNode(internal.Left, writer)
		writeNode(internal.Right, writer)
	}
}

func PrettyPrintSerializedTree(data []byte) {
	for _, b := range data {
		switch b {
		case 0:
			fmt.Print("0 ")
		case 1:
			fmt.Print("1 ")
		default:
			fmt.Printf("%c ", b)
		}
	}
	fmt.Println()
}
