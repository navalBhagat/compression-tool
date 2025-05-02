package main

import (
	"bufio"
	"compression-tool/pkg/frequencyMap"
	"compression-tool/pkg/huffmanEncodingTree"
	"compression-tool/pkg/standardFunctions"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	var scanner *bufio.Scanner
	switch len(args) {
	case 1:
		filename := args[0]
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("Unable to read file: %v", err)
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	default:
		if standardFunctions.IsInputFromPipe() {
			scanner = bufio.NewScanner(os.Stdin)
		} else {
			standardFunctions.PrintUsageAndExit()
		}
	}

	frequencyMap, err := frequencyMap.CreateFrequencyMap(scanner)
	if err != nil {
		fmt.Print("Couldn't calculate frequency map")
	}
	huffTree := huffmanEncodingTree.BuildTree(frequencyMap)
	prefixTable := huffmanEncodingTree.PrefixTable(huffTree)
	fmt.Println(prefixTable)
}
