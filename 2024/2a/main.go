package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

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

		var levels []int
		prev := 0
		for n, c := range line {
			if c == ' ' || n == len(line)-1 {
				last := n
				if n == len(line)-1 {
					last++
				}
				l, err := strconv.Atoi(line[prev:last])
				if err != nil {
					log.Fatalf("could not parse level %d, from %d to %d on line %s: %v", len(levels), prev, n, line, err)
				}
				levels = append(levels, l)
				prev = n + 1
			}
		}

		if len(levels) < 2 {
			log.Fatalf("line %s has less than two levels", line)
		}
		increasing := false
		if levels[0] < levels[1] {
			increasing = true
		}
		prev = levels[0]
		valid := true
		for n := 1; n < len(levels); n++ {
			if increasing {
				if levels[n] <= prev || levels[n]-prev > 3 {
					valid = false
					break
				}
			} else {
				if levels[n] >= prev || prev-levels[n] > 3 {
					valid = false
					break
				}
			}
			prev = levels[n]
		}
		if !valid {
			continue
		}

		fmt.Println(levels, increasing)
		sum++
	}

	fmt.Println("Answer:", sum)
}
