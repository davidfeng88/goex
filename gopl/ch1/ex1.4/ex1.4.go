// Based on dup2

// Prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
// If there are duplicated lines, the filename is also included.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	var fileNamePrinted bool
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if counts[line] > 1 && !fileNamePrinted {
			fmt.Printf("The file %s contains duplicated lines:\n", f.Name())
			// if f is stdin, f.Name() returns "/dev/stdin"
			fileNamePrinted = true
		}

	}
	// NOTE: ignoring potential errors from input.Err()
}
