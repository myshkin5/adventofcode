package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/myshkin5/adventofcode/paths"
)

func main() {
	f, err := os.Open(filepath.Join(paths.SourcePath(), "input.txt"))
	if err != nil {
		log.Fatalf("could not open file: %#v", err)
	}
	s := bufio.NewScanner(f)
	var left []int
	right := make(map[int]int)
	for s.Scan() {
		line := string(s.Bytes())
		space := strings.Index(line, " ")
		if space == -1 {
			log.Fatalf("could not parse line: %#v", line)
		}
		l, err := strconv.Atoi(line[:space])
		if err != nil {
			log.Fatalf("could not parse left of line: %#v: %#v", line, err)
		}
		r, err := strconv.Atoi(line[space+3:])
		if err != nil {
			log.Fatalf("could not parse right of line: %#v: %#v", line, err)
		}
		left = append(left, l)
		right[r] = right[r] + 1
	}

	sum := 0
	for _, l := range left {
		sum += l * right[l]
	}

	fmt.Println("Answer:", sum)
}
