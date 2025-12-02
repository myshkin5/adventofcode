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

	rules := make(map[int]map[int]struct{})

	for s.Scan() {
		line := string(s.Bytes())
		if len(line) == 0 {
			break
		}

		splits := strings.Split(line, "|")
		if len(splits) != 2 {
			log.Fatalf("invalid line: %#v", line)
		}
		l, err := strconv.Atoi(splits[0])
		if err != nil {
			log.Fatalf("invalid line: %#v", line)
		}
		r, err := strconv.Atoi(splits[1])
		if err != nil {
			log.Fatalf("invalid line: %#v", line)
		}
		rule, ok := rules[l]
		if !ok {
			rule = make(map[int]struct{})
		}
		rule[r] = struct{}{}
		rules[l] = rule
	}

	sum := 0
	for s.Scan() {
		line := string(s.Bytes())

		splits := strings.Split(line, ",")
		if len(splits)%2 != 1 {
			log.Fatalf("invalid line: %#v", line)
		}

		inOrder := true
		var prevs []int
		for _, v := range splits {
			n, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf("invalid line: %#v", line)
			}

			rule := rules[n]
			for _, p := range prevs {
				if _, ok := rule[p]; ok {
					inOrder = false
					break
				}
			}

			if !inOrder {
				break
			}
			prevs = append(prevs, n)
		}

		if inOrder {
			fmt.Println("In order", line)
			sum += prevs[len(prevs)/2]
		}
	}

	fmt.Println("Answer:", sum)
}
