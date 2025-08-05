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
	case 2:
		filename := args[0]
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("Unable to read file: %v", err)
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	case 1:
		if standardFunctions.IsInputFromPipe() {
			scanner = bufio.NewScanner(os.Stdin)
		} else {
			standardFunctions.PrintUsageAndExit()
		}
	default:
		standardFunctions.PrintUsageAndExit()
	}

	frequencyMap, err := frequencyMap.CreateFrequencyMap(scanner)
	if err != nil {
		fmt.Print("Couldn't calculate frequency map")
	}
	huffTree := huffmanEncodingTree.BuildTree(frequencyMap)
	huffmanEncodingTree.PrettyPrintTree(huffTree.RootNode(), "")
	prefixTable := huffmanEncodingTree.PrefixTable(huffTree)
	fmt.Println(prefixTable)
	serializedTree := huffmanEncodingTree.SerializeTree(huffTree)

	outputFile := standardFunctions.DetermineOutputFilename()
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer outFile.Close()

	outFile.Write(serializedTree)
	fmt.Printf("File compressed and written to %s\n", outputFile)
}
