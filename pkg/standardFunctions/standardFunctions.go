package standardFunctions

import (
	"fmt"
	"os"
)

func IsInputFromPipe() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func PrintUsageAndExit() {
	fmt.Println("Usage: compression-tool <filename> <output filename> or cat <filename> | compression-tool <output filename>")
	os.Exit(1)
}

func DetermineOutputFilename() string {
	if len(os.Args) > 2 {
		return os.Args[2]
	}
	return "output.huff"
}

func MapsEqual[K comparable, V comparable](a, b map[K]V) bool {
	if len(a) != len(b) {
		return false
	}
	for key, valA := range a {
		valB, ok := b[key]
		if !ok || valA != valB {
			return false
		}
	}
	return true
}
