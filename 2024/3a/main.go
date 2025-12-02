package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/myshkin5/adventofcode/paths"
)

func main() {
	f, err := os.Open(filepath.Join(paths.SourcePath(), "input.txt"))
	if err != nil {
		log.Fatalf("could not open file: %#v", err)
	}
	s := bufio.NewScanner(f)

	rexp := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)

	sum := 0

	var line string
	for s.Scan() {
		line = string(s.Bytes())

		muls := rexp.FindAllString(line, -1)
		for _, m := range muls {
			groups := rexp.FindStringSubmatch(m)
			if len(groups) != 3 {
				log.Fatalf("malformed line: %#v", m)
			}
			x, err := strconv.Atoi(groups[1])
			if err != nil {
				log.Fatalf("malformed line: %#v", m)
			}
			y, err := strconv.Atoi(groups[2])
			if err != nil {
				log.Fatalf("malformed line: %#v", m)
			}
			sum += x * y
		}
	}

	fmt.Println("Answer:", sum)
}
