package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	use := flag.String("count", "words", "Indicate what to count: "+
		"words, lines, bytes, runes.")
	flag.Parse()
	fmt.Println(count(os.Stdin, *use))
}

func count(r io.Reader, countWhat string) int {
	// Scan given reader by words
	scanner := bufio.NewScanner(r)

	// Determine what to count
	switch countWhat {
	case "words":
		scanner.Split(bufio.ScanWords)
	case "lines":
		scanner.Split(bufio.ScanLines)
	case "bytes":
		scanner.Split(bufio.ScanBytes)
	case "runes":
		scanner.Split(bufio.ScanRunes)
	default:
		fmt.Print("unknown flag value, usage:\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Count words
	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}
