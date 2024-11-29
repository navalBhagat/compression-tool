package main

import (
	"bufio"
	"compression-tool/pkg/frequencyMap"
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
		if isInputFromPipe() {
			scanner = bufio.NewScanner(os.Stdin)
		} else {
			printUsageAndExit()
		}
	}

	frequencyMap, err := frequencyMap.CreateFrequencyMap(scanner)
	if err != nil {
		fmt.Print("Couldn't calculate frequency map")
	}
	fmt.Println(frequencyMap["X"], frequencyMap["t"])
}

func isInputFromPipe() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func printUsageAndExit() {
	fmt.Println("Usage: compression-tool <filename> or cat <filename> | compression-tool")
	os.Exit(1)
}
