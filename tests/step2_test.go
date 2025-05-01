package tests

import (
	"bufio"
	"os"
	"testing"

	"compression-tool/pkg/frequencyMap"
	"compression-tool/pkg/huffmanEncodingTree"
)

var expectedTree = huffmanEncodingTree.HuffTree{
	Root: &huffmanEncodingTree.HuffInternalNode{
		NodeWeight: 306,
		Left: &huffmanEncodingTree.HuffLeafNode{
			Element:    "E",
			NodeWeight: 120,
		},
		Right: &huffmanEncodingTree.HuffInternalNode{
			NodeWeight: 186,
			Left: &huffmanEncodingTree.HuffInternalNode{
				NodeWeight: 79,
				Left: &huffmanEncodingTree.HuffLeafNode{
					Element:    "U",
					NodeWeight: 37,
				},
				Right: &huffmanEncodingTree.HuffLeafNode{
					Element:    "D",
					NodeWeight: 42,
				},
			},
			Right: &huffmanEncodingTree.HuffInternalNode{
				NodeWeight: 107,
				Left: &huffmanEncodingTree.HuffLeafNode{
					Element:    "L",
					NodeWeight: 42,
				},
				Right: huffmanEncodingTree.HuffInternalNode{
					NodeWeight: 65,
					Left: &huffmanEncodingTree.HuffLeafNode{
						Element:    "C",
						NodeWeight: 32,
					},
					Right: &huffmanEncodingTree.HuffInternalNode{
						NodeWeight: 33,
						Left: &huffmanEncodingTree.HuffInternalNode{
							NodeWeight: 9,
							Left: &huffmanEncodingTree.HuffLeafNode{
								Element:    "Z",
								NodeWeight: 2,
							},
							Right: &huffmanEncodingTree.HuffLeafNode{
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

var filename = "./testdata/step2_test.txt"

func testHuffmanTreeFromFile(t *testing.T) {
	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Unable to read file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	freqMap, _ := frequencyMap.CreateFrequencyMap(scanner)
	huffTree := huffmanEncodingTree.BuildTree(freqMap)

	if huffTree != expectedTree {
		t.Errorf("Expected to have map %s, but got %s", expectedTree, huffTree)
	}
}

func testHuffmanTreeFromStdIn(t *testing.T) {
	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	origStdin := os.Stdin
	defer func() { os.Stdin = origStdin }()
	os.Stdin = file

	scanner := bufio.NewScanner(os.Stdin)
	freqMap, _ := frequencyMap.CreateFrequencyMap(scanner)
	huffTree := huffmanEncodingTree.BuildTree(freqMap)

	if huffTree != expectedTree {
		t.Errorf("Expected to have map %s, but got %s", expectedTree, huffTree)
	}
}
