package tests

import (
	"bufio"
	"compression-tool/pkg/frequencyMap"
	"compression-tool/pkg/huffmanEncodingTree"
	"compression-tool/pkg/standardFunctions"
	"os"
	"testing"
)

var expectedPrefixTable = map[string]string{
	"E": "0",
	"U": "100",
	"D": "101",
	"L": "110",
	"C": "1110",
	"Z": "111100",
	"K": "111101",
	"M": "11111",
}

func testPrefixTableFromFile(t *testing.T) {
	filename := "./testdata/test_small.txt"
	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Unable to read file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	freqMap, _ := frequencyMap.CreateFrequencyMap(scanner)
	huffTree := huffmanEncodingTree.BuildTree(freqMap)
	prefixTable := huffmanEncodingTree.PrefixTable(huffTree)

	if !standardFunctions.MapsEqual(prefixTable, expectedPrefixTable) {
		t.Errorf("Expected to have map %s, but got %s", expectedPrefixTable, prefixTable)
	}
}

func testPrefixTableFromStdIn(t *testing.T) {
	filename := "./testdata/test_small.txt"
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
	prefixTable := huffmanEncodingTree.PrefixTable(huffTree)

	if !standardFunctions.MapsEqual(prefixTable, expectedPrefixTable) {
		t.Errorf("Expected to have map %s, but got %s", expectedPrefixTable, prefixTable)
	}
}
