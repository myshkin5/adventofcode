package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"unicode"

	"github.com/myshkin5/adventofcode/paths"
	"github.com/myshkin5/adventofcode/strs"
)

var rexp = regexp.MustCompile("[0-9]+")

func main() {
	f, err := os.Open(filepath.Join(paths.SourcePath(), "input.txt"))
	if err != nil {
		log.Fatalf("could not open file: %#v", err)
	}
	s := bufio.NewScanner(f)
	var sum int
	previous := ""
	line := ""
	if !s.Scan() {
		log.Fatalf("no lines found in input")
	}
	next := string(s.Bytes())
	for s.Scan() {
		previous = line
		line = next
		next = string(s.Bytes())
		sum += checkAll(previous, line, next)
	}
	previous = line
	line = next
	next = ""
	sum += checkAll(previous, line, next)

	fmt.Println("Answer:", sum)
}

func checkAll(previous, line, next string) int {
	sum := 0
	for _, indices := range rexp.FindAllStringIndex(line, -1) {
		start := indices[0]
		end := indices[1]
		partNoS := line[start:end]

		if start > 0 {
			start--
		}
		if end < len(line) {
			end++
		}

		isSymbol := func(r rune) bool {
			return !unicode.IsDigit(r) && r != '.'
		}

		checkLine := func(nOrP string) bool {
			for n := 0; n < len(nOrP); n++ {
				if isSymbol(rune(nOrP[n])) {
					return true
				}
			}
			return false
		}

		prevCheck := previous != "" && checkLine(previous[start:end])
		nextCheck := next != "" && checkLine(next[start:end])
		before := isSymbol(rune(line[start]))
		after := isSymbol(rune(line[end-1]))

		log.Println(partNoS, prevCheck, nextCheck, before, after)

		if prevCheck || nextCheck || before || after {
			sum += strs.Atoi(partNoS)
		}
	}

	return sum
}
