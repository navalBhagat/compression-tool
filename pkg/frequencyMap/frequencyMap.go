package frequencyMap

import (
	"bufio"
	"fmt"
	"unicode/utf8"
)

func CreateFrequencyMap(scanner *bufio.Scanner) (map[string]int, error) {
	freq := make(map[string]int)

	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		letter, _ := utf8.DecodeRune(scanner.Bytes())
		if letter == utf8.RuneError {
			return nil, fmt.Errorf("Invalid utf8 encoding")
		}
		freq[string(letter)]++
	}

	return freq, nil
}
