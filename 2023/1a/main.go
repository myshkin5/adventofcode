package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"unicode"

	"github.com/myshkin5/adventofcode/paths"
)

func main() {
	f, err := os.Open(filepath.Join(paths.SourcePath(), "input.txt"))
	if err != nil {
		log.Fatalf("could not open file: %#v", err)
	}
	s := bufio.NewScanner(f)
	sum := 0
	for s.Scan() {
		line := string(s.Bytes())
		foundFirst := false
		var first, last int
		for _, c := range line {
			if !unicode.IsDigit(c) {
				continue
			}

			last = int(c - '0')
			if !foundFirst {
				foundFirst = true
				first = last
			}
		}
		sum += first*10 + last
		fmt.Println(first*10+last, line)
	}

	fmt.Println("Answer:", sum)
}
