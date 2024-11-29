package tests

import (
	"bufio"
	"compression-tool/pkg/frequencyMap"
	"os"
	"testing"
)

func TestCountFromFile(t *testing.T) {
	filename := "./testdata/test.txt"
	expectedX := 333
	expectedt := 223000

	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Unable to read file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result, err := frequencyMap.CreateFrequencyMap(scanner)

	if err != nil {
		t.Errorf("Couldn't generate frequency map")
	}
	if result["X"] != expectedX {
		t.Errorf("Expected \"X\" to have frequency %d, but got %d", expectedX, result["X"])
	}

	if result["t"] != expectedt {
		t.Errorf("Expected \"t\" to have frequency %d, but got %d", expectedt, result["t"])
	}
}

func TestCountFromStdIn(t *testing.T) {
	expectedX := 333
	expectedt := 223000

	// Simulate input from stdin by opening the test file and redirecting stdin
	file, err := os.Open("./testdata/test.txt")
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	origStdin := os.Stdin
	defer func() { os.Stdin = origStdin }()
	os.Stdin = file

	scanner := bufio.NewScanner(os.Stdin)
	result, err := frequencyMap.CreateFrequencyMap(scanner)

	if err != nil {
		t.Errorf("Couldn't generate frequency map")
	}
	if result["X"] != expectedX {
		t.Errorf("Expected \"X\" to have frequency %d, but got %d", expectedX, result["X"])
	}

	if result["t"] != expectedt {
		t.Errorf("Expected \"t\" to have frequency %d, but got %d", expectedt, result["t"])
	}
}
