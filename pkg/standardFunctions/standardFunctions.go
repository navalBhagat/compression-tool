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
	fmt.Println("Usage: compression-tool <filename> or cat <filename> | compression-tool")
	os.Exit(1)
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
