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

	var line string
	for s.Scan() {
		line += string(s.Bytes())
	}

	dosDonts := regexp.MustCompile(`do\(\)|don't\(\)`).FindAllIndex([]byte(line), -1)
	dos := []string{line[0:dosDonts[0][0]]}
	do := true
	for n, d := range dosDonts {
		doDont := line[d[0]:d[1]]
		switch {
		case doDont == "do()":
			do = true
		case doDont == "don't()":
			do = false
		}
		if do {
			end := len(line)
			if n < len(dosDonts)-1 {
				end = dosDonts[n+1][0]
			}
			dos = append(dos, line[d[1]:end])
		}
	}

	rexp := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	sum := 0
	for _, d := range dos {
		muls := rexp.FindAllString(d, -1)

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
