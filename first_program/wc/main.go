package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(count(os.Stdin))
}

func count(r io.Reader) int {
	// Scan given reader by words
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	// Count words
	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}
