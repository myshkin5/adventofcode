package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/myshkin5/adventofcode/paths"
	"github.com/myshkin5/adventofcode/strs"
)

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
	gStart := 0
	count := 0
	for {
		gIdx := indexAt(line, "*", gStart)
		if gIdx == -1 {
			break
		}
		gStart = gIdx + 2

		partNos := findLinePartNos(previous, gIdx)
		partNos = append(partNos, findLinePartNos(next, gIdx)...)
		partNos = append(partNos, findLinePartNos(line, gIdx)...)

		if len(partNos) > 2 {
			log.Fatalf("found more than two parts nos for gear %d on line %s", gIdx, line)
		}
		if len(partNos) == 2 {
			count += partNos[0] * partNos[1]
		}
	}

	return count
}

func findLinePartNos(line string, gIdx int) []int {
	var partNos []int
	if unicode.IsDigit(rune(line[gIdx])) {
		start := findFirstDigit(line, gIdx)
		end := findLastDigit(line, gIdx)
		partNos = append(partNos, strs.Atoi(line[start:end]))
	} else {
		if gIdx > 0 && unicode.IsDigit(rune(line[gIdx-1])) {
			start := findFirstDigit(line, gIdx-1)
			end := gIdx
			partNos = append(partNos, strs.Atoi(line[start:end]))
		}
		if gIdx < len(line) && unicode.IsDigit(rune(line[gIdx+1])) {
			start := gIdx + 1
			end := findLastDigit(line, gIdx+1)
			partNos = append(partNos, strs.Atoi(line[start:end]))
		}
	}
	return partNos
}

func findFirstDigit(line string, idx int) int {
	for n := idx; n >= 0; n-- {
		if !unicode.IsDigit(rune(line[n])) {
			return n + 1
		}
	}
	return 0
}

func findLastDigit(line string, idx int) int {
	for n := idx; n < len(line); n++ {
		if !unicode.IsDigit(rune(line[n])) {
			return n
		}
	}
	return len(line)
}

func indexAt(s, sep string, n int) int {
	idx := strings.Index(s[n:], sep)
	if idx > -1 {
		idx += n
	}
	return idx
}
