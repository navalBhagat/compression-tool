package tests

import (
	"bytes"
	"compression-tool/pkg/huffmanEncodingTree"
	"testing"
)

var huffTree = huffmanEncodingTree.HuffTree{
	Root: huffmanEncodingTree.HuffInternalNode{
		NodeWeight: 306,
		Left: huffmanEncodingTree.HuffLeafNode{
			Element:    "E",
			NodeWeight: 120,
		},
		Right: huffmanEncodingTree.HuffInternalNode{
			NodeWeight: 186,
			Left: huffmanEncodingTree.HuffInternalNode{
				NodeWeight: 79,
				Left: huffmanEncodingTree.HuffLeafNode{
					Element:    "U",
					NodeWeight: 37,
				},
				Right: huffmanEncodingTree.HuffLeafNode{
					Element:    "D",
					NodeWeight: 42,
				},
			},
			Right: huffmanEncodingTree.HuffInternalNode{
				NodeWeight: 107,
				Left: huffmanEncodingTree.HuffLeafNode{
					Element:    "L",
					NodeWeight: 42,
				},
				Right: huffmanEncodingTree.HuffInternalNode{
					NodeWeight: 65,
					Left: huffmanEncodingTree.HuffLeafNode{
						Element:    "C",
						NodeWeight: 32,
					},
					Right: huffmanEncodingTree.HuffInternalNode{
						NodeWeight: 33,
						Left: huffmanEncodingTree.HuffInternalNode{
							NodeWeight: 9,
							Left: huffmanEncodingTree.HuffLeafNode{
								Element:    "Z",
								NodeWeight: 2,
							},
							Right: huffmanEncodingTree.HuffLeafNode{
								Element:    "K",
								NodeWeight: 7,
							},
						},
						Right: huffmanEncodingTree.HuffLeafNode{
							Element:    "M",
							NodeWeight: 24,
						},
					},
				},
			},
		},
	},
}

func TestSerializeTree(t *testing.T) {
	result := huffmanEncodingTree.SerializeTree(huffTree)
	expected := []byte{
		0,      // root internal
		1, 'E', // left leaf: E
		0,      // right internal
		0,      // left internal of right
		1, 'U', // left leaf: U
		1, 'D', // right leaf: D
		0,      // right internal of right
		1, 'L', // left leaf: L
		0,      // right internal of right-right
		1, 'C', // left leaf: C
		0,      // right internal of right-right-right
		0,      // left internal of above
		1, 'Z', // left leaf: Z
		1, 'K', // right leaf: K
		1, 'M', // right leaf: M
	}
	if !bytes.Equal(result, expected) {
		t.Errorf("SerializeTree() = %v, want %v", result, expected)
	}
}
